package authsrv

import (
	"context"
	"database/sql"
	"errors"

	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUser - method for retrieving user information based on JWT.
func (s *AuthServer) GetUser(ctx context.Context, in *ssov1.GetUserRequest) (*ssov1.GetUserResponse, error) {
	log := s.Log

	userID, err := s.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	// Get info about user from db
	user, err := s.Queries.GetUserByID(ctx, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.NotFound, "user not found")
	} else if err != nil {
		log.Errorf("error getting user: %v", err)
		return nil, status.Error(codes.Internal, "failed to retrieve user")
	}

	// Inf about user
	return &ssov1.GetUserResponse{
		UserId:   user.ID.String(),
		Username: user.Username.String,
		Email:    user.Email.String,
	}, nil
}
