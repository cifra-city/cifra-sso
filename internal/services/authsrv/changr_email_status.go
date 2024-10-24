package authsrv

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cifra-city/cifra-sso/internal/db/data"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *AuthServer) ChangeEmailStatus(ctx context.Context, in *ssov1.ChangeEmailStatusRequest) (*ssov1.Empty, error) {
	log := s.Log

	userID, err := s.Authenticate(ctx)

	user, err := s.Queries.GetUserByID(ctx, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.NotFound, "user not found")
	} else if err != nil {
		log.Errorf("error getting user: %v", err)
		return nil, status.Error(codes.Internal, "failed to retrieve user")
	}

	stmt := data.UpdateEmailStatusByIDParams{
		ID:          userID,
		EmailStatus: in.EmailStatus,
	}

	_, err = s.Queries.UpdateEmailStatusByID(ctx, stmt)
	if err != nil {
		log.Errorf("error updating email status user %s: %v", user.Username, err)
		return nil, status.Error(codes.Internal, "Error updating email status user")
	}

	log.Infof("email updated for user: %s - status %v", user.Username, in.EmailStatus)
	return &ssov1.Empty{}, nil
}
