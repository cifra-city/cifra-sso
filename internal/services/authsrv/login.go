package authsrv

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/domain"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Login - method for user login and JWT token generation.
func (s *AuthServer) Login(ctx context.Context, in *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	// Convert input strings to sql.NullString.
	email := domain.ToNullString(in.Email)
	username := domain.ToNullString(in.Username)
	password := domain.ToNullString(in.Password)

	// Validate the incoming request data.
	if !email.Valid && !username.Valid {
		return nil, status.Error(codes.InvalidArgument, "email or username is required")
	}
	if !password.Valid {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	// Fetch the user from the database.
	var user data.User
	var err error
	if email.Valid {
		user, err = s.Queries.GetUserByEmail(ctx, email)
	} else {
		user, err = s.Queries.GetUserByUsername(ctx, username)
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.Unauthenticated, "user not found")
	} else if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve user")
	}

	// Compare the provided password with the stored hashed password.
	err = bcrypt.CompareHashAndPassword([]byte(user.PasHash.String), []byte(password.String))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid password")
	}

	// Generate a JWT token.
	token, err := s.generateJWT(user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	// Return the token in the response.
	return &ssov1.LoginResponse{Token: token}, nil
}

// generateJWT creates a JWT token for the authenticated user.
func (s *AuthServer) generateJWT(userID uuid.UUID) (string, error) {
	// Define token expiration time.
	expirationTime := time.Now().Add(s.Config.JWT.TokenLifetime * time.Second)

	// Create the JWT claims, which include the user ID and expiration time.
	claims := &jwt.RegisteredClaims{
		Subject:   userID.String(), // Use user ID as the subject in string format.
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	// Create the JWT token using the claims and the secret key.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.Config.JWT.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
