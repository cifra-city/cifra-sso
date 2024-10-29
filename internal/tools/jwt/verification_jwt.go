package jwt

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// VerificationJWT validates the JWT token and returns the user ID if valid.
func VerificationJWT(ctx context.Context, log *logrus.Logger, sk string) (uuid.UUID, error) {
	tokenString, err := ExtractToken(ctx)
	if err != nil {
		log.Infof("error extracting token: %v", err)
		return uuid.Nil, status.Error(codes.Internal, "Error extracting token")
	}

	// Parse and validate the JWT token.
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(sk), nil
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
