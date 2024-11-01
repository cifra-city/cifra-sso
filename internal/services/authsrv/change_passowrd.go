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

func (s *AuthServer) ChangePassword(ctx context.Context, in *ssov1.ChangePasswordRequest) (*ssov1.Empty, error) {
	log := s.Log

	userID, err := s.Authenticate(ctx)

	user, err := s.Queries.GetUserByID(ctx, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.NotFound, "user not found")
	} else if err != nil {
		log.Errorf("error getting user: %v", err)
		return nil, status.Error(codes.Internal, "failed to retrieve user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(in.OldPassword))
	if err != nil {
		log.Debugf("incorect password for user: %s", user.Username)
		return nil, status.Error(codes.Unauthenticated, "invalid password")
	}

	newpassword, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("error generating new password for user: %s", user.Username)
		return nil, status.Error(codes.Internal, "Server Error")
	}

	//TODO add auth email

	stmt := data.UpdatePasswordByIDParams{
		ID:       userID,
		PassHash: string(newpassword),
	}

	_, err = s.Queries.UpdatePasswordByID(ctx, stmt)
	if err != nil {
		log.Errorf("error updating password user %s: %v", user.Username, err)
		return nil, status.Error(codes.Internal, "Server Error")
	}

	log.Infof("password updated for user: %s", user.Username)
	return &ssov1.Empty{}, nil
}
