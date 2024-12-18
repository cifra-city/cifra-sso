package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cifra-city/cifra-jwt"
	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/domain/entities"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ChangeUsername(ctx context.Context, in *ssov1.ChangeUsernameReq) (*ssov1.Empty, error) {
	log := s.Log

	userID, err := cifra_jwt.VerificationJWT(ctx, log, s.Config.JWT.SecretKey)
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

	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(in.Password))
	if err != nil {
		log.Debugf("incorect password for user: %s", user.Username)
		return nil, status.Error(codes.Unauthenticated, "invalid password")
	}

	eve, err := s.ActionPermission.GetEvent(user.Username)
	if err != nil {
		log.Errorf("error checking in queue: %v", err)
		return nil, status.Error(codes.Internal, "failed to check in queue")
	}
	if eve != entities.ChangeUsername {
		log.Errorf("user %s is not in the change username queue", user.Username)
		return nil, status.Error(codes.PermissionDenied, "user is not in the change username queue")
	}

	stmt := data.UpdateUsernameByIDParams{
		Username: in.NewUsername,
		ID:       userID,
	}

	_, err = s.Queries.UpdateUsernameByID(ctx, stmt)
	if err != nil {
		log.Errorf("error updating user: %v", err)
		return nil, status.Error(codes.Internal, "failed to update user")
	}

	log.Infof("user %s has change username to %s", user.Username, in.NewUsername)
	return &ssov1.Empty{}, nil
}
