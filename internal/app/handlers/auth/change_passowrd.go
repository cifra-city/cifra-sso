package auth

import (
	"context"
	"database/sql"
	"errors"

	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/domain/entities"
	"github.com/cifra-city/cifra-sso/internal/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ChangePass(ctx context.Context, in *ssov1.ChangePassReq) (*ssov1.Empty, error) {
	log := s.Log

	userID, err := jwt.VerificationJWT(ctx, log, s.Config.JWT.SecretKey)
	if err != nil {
		log.Error("Error getting user from JWT-token: %s", err)
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	user, err := s.Queries.GetUserByID(ctx, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.NotFound, "user not found")
	} else if err != nil {
		log.Errorf("error getting user: %v", err)
		return nil, status.Error(codes.Internal, "failed to retrieve user")
	}

	if !user.EmailStatus {
		log.Errorf("user %s has not confirmed their email", user.Username)
		return nil, status.Error(codes.PermissionDenied, "user has not confirmed your email")
	}

	eve, err := s.ActionPermission.GetEvent(user.Username)
	if err != nil {
		log.Errorf("error checking in queue: %v", err)
		return nil, status.Error(codes.Internal, "failed to check in queue")
	}
	if eve != entities.ChangePassword {
		log.Errorf("user %s is not in the change passwoed queue", user.Username)
		return nil, status.Error(codes.PermissionDenied, "user is not in the change email password")
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("error generating new password for user: %s", user.Username)
		return nil, status.Error(codes.Internal, "Server Error")
	}

	stmt := data.UpdatePasswordByIDParams{
		ID:       userID,
		PassHash: string(newPassword),
	}

	_, err = s.Queries.UpdatePasswordByID(ctx, stmt)
	if err != nil {
		log.Errorf("error updating password user %s: %v", user.Username, err)
		return nil, status.Error(codes.Internal, "Server Error")
	}

	log.Infof("password updated for user: %s", user.Username)
	return &ssov1.Empty{}, nil
}
