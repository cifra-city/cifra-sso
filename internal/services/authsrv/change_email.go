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

func (s *AuthServer) ChangeEmail(ctx context.Context, in *ssov1.ChangeEmailRequest) (*ssov1.Empty, error) {
	log := s.Log

	userID, err := s.Authenticate(ctx)

	user, err := s.Queries.GetUserByID(ctx, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.NotFound, "user not found")
	} else if err != nil {
		log.Errorf("error getting user: %v", err)
		return nil, status.Error(codes.Internal, "failed to retrieve user")
	}

	userEmail, err := s.Queries.GetUserByEmail(ctx, user.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.NotFound, "user not found")
	} else if err != nil {
		log.Errorf("error getting user: %v", err)
		return nil, status.Error(codes.Internal, "failed to retrieve user")
	}

	if user != userEmail {
		log.Errorf("user id and email do not match data in request: %s, %s - in db: %s %s",
			in.OldEmail, userID, user.Email, user.ID)
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	//TODO add auth email

	stmt := data.UpdateEmailByIDParams{
		ID:    userID,
		Email: in.NewEmail,
	}

	_, err = s.Queries.UpdateEmailByID(ctx, stmt)
	if err != nil {
		log.Errorf("error updating email user %s: %v", user.Username, err)
		return nil, status.Error(codes.Internal, "Error updating email for user")
	}

	log.Infof("email updated for user: %s - new email %v", in.OldEmail, in.NewEmail)
	return &ssov1.Empty{}, nil
}
