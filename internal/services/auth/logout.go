package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cifra-city/cifra-sso/internal/tools/jwt"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Logout - method for user logout, clears the JWT cookie.
func (s *AuthServer) Logout(ctx context.Context, in *ssov1.Empty) (*ssov1.Empty, error) {

	log := s.Log

	userID, err := jwt.VerificationJWT(ctx, log, s.Config.JWT.SecretKey)
	if err != nil {
		log.Error("Error getting user from JWT-token: %s", err)
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	_, err = s.Queries.GetUserByID(ctx, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.NotFound, "user not found")
	} else if err != nil {
		log.Errorf("error getting user: %v", err)
		return nil, status.Error(codes.Internal, "failed to retrieve user")
	}

	log.Infof("user with %s success logout fron device", in.String())

	return &ssov1.Empty{}, nil // Clear the JWT cookie at frontend
}
