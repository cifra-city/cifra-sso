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

	log := s.Log
	// Validate the incoming request data.
	if !email.Valid && !username.Valid {
		log.Debugf("email or username are missing")
		return nil, status.Error(codes.InvalidArgument, "email or username is required")
	}
	if !password.Valid {
		log.Debugf("password is missing")
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	// Fetch the user from the database.
	var user data.User
	var err error
	if email.Valid && username.Valid {
		user, err = s.Queries.GetUserByEmail(ctx, email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, status.Error(codes.NotFound, "user not found")
			}
			log.Errorf("Error %s", err)
			return nil, status.Error(codes.Internal, "error getting user")
		}
		if user, err = s.Queries.GetUserByUsername(ctx, username); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, status.Error(codes.NotFound, "user not found")
			}
			log.Errorf("error is %s", err)
			return nil, status.Error(codes.Internal, "error getting user")
		}
	}

	//TODO add logic for proof email by sending list with code

	if errors.Is(err, sql.ErrNoRows) {
		log.Debugf("user not found %s", username.String)
		return nil, status.Error(codes.Unauthenticated, "user not found")
	} else if err != nil {
		log.Debugf("error getting user %s", err)
		return nil, status.Error(codes.Internal, "failed to retrieve user")
	}

	// Compare the provided password with the stored hashed password.
	err = bcrypt.CompareHashAndPassword([]byte(user.PasHash.String), []byte(password.String))
	if err != nil {
		log.Debugf("incorect password for user: %s", user.Username.String)
		return nil, status.Error(codes.Unauthenticated, "invalid password")
	}

	// Generate a JWT token.
	token, err := s.generateJWT(user.ID)
	if err != nil {
		log.Error("error creating jwt token %s", err)
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	// Return the token in the response.
	log.Infof("successfully logged in user %s", user.Username.String)
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
