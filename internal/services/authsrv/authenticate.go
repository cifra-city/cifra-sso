package authsrv

import (
	"context"

	"github.com/cifra-city/cifra-sso/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Authenticate validates the JWT token and returns the user ID if valid.
func (s *AuthServer) Authenticate(ctx context.Context) (uuid.UUID, error) {
	log := s.Log
	tokenString, err := domain.ExtractToken(ctx)
	if err != nil {
		log.Infof("error extracting token: %v", err)
		return uuid.Nil, status.Error(codes.Internal, "Error extracting token")
	}

	// Parse and validate the JWT token.
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config.JWT.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return uuid.Nil, status.Error(codes.Unauthenticated, "invalid or expired token")
	}

	// Extract the user ID from the claims (assuming it's stored in the "Subject" field).
	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.Nil, status.Error(codes.Internal, "failed to parse user ID")
	}

	return userID, nil
}
