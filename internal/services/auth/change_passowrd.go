package authsrv

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cifra-city/cifra-sso/internal/db/data"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *AuthServer) ChangePassword(ctx context.Context, in *ssov1.ChangePassReq) (*ssov1.Empty, error) {
	log := s.Log

	userID, err := s.Authenticate(ctx)
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

	eve, err := s.Events.GetEvent(user.Username)
	if err != nil {
		log.Errorf("error checking in queue: %v", err)
		return nil, status.Error(codes.Internal, "failed to check in queue")
	}
	if eve != CHANGE_PASS {
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
