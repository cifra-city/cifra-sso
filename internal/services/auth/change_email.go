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

func (s *AuthServer) ChangeEmail(ctx context.Context, in *ssov1.ChangeEmailReq) (*ssov1.Empty, error) {
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
	if eve != CHANGE_EMAIL {
		log.Errorf("user %s is not in the change email queue", user.Username)
		return nil, status.Error(codes.PermissionDenied, "user is not in the change email queue")
	}

	stmt := data.UpdateEmailByIDParams{
		ID:          userID,
		Email:       in.NewEmail,
		EmailStatus: false,
	}

	_, err = s.Queries.UpdateEmailByID(ctx, stmt)
	if err != nil {
		log.Errorf("error updating email user %s: %v", user.Username, err)
		return nil, status.Error(codes.Internal, "Error updating email for user")
	}

	log.Infof("email updated for user: %s - new email %v", user.Username, in.NewEmail)
	return &ssov1.Empty{}, nil
}
