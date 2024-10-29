package authsrv

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/cifra-city/cifra-sso/internal/db/data"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Login - method for user login and JWT token generation.
func (s *AuthServer) Login(ctx context.Context, in *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	log := s.Log

	// Fetch the user from the database.
	var user data.UsersSecret
	var err error

	if in.Username != nil {
		user, err = s.Queries.GetUserByUsername(ctx, *in.Username)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, status.Error(codes.NotFound, "user not found")
			}
			log.Errorf("Error %s", err)
			return nil, status.Error(codes.Internal, "error getting user")
		}
		log.Infof("Start try to response by %s", user.Username)
	} else if in.Email != nil {
		user, err = s.Queries.GetUserByEmail(ctx, *in.Email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, status.Error(codes.NotFound, "user not found")
			}
			log.Errorf("Error %s", err)
			return nil, status.Error(codes.Internal, "error getting user")
		}
		log.Infof("Start try to response by %s", user.Email)
	} else {
		return nil, status.Error(codes.InvalidArgument, "username or email not provided")
	}

	if user.EmailStatus {
		if !s.Email.CheckInEmailList(user.Username, in.Code) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
	}

	// Compare the provided password with the stored hashed password.
	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(in.Password))
	if err != nil {
		log.Debugf("incorect password for user: %s", user.Username)
		return nil, status.Error(codes.Unauthenticated, "invalid password")
	}

	// Generate a JWT token.
	token, err := s.generateJWT(user.ID)
	if err != nil {
		log.Error("error creating jwt token %s", err)
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	// Return the token in the response.
	log.Infof("successfully logged in user %s", user.Username)
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
