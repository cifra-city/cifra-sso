package authsrv

import (
	"context"
	"database/sql"

	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/domain"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register - method for registering a new user.
func (s *AuthServer) Register(ctx context.Context, in *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	// Validate the incoming request db.
	if in.Email == "" || in.Username == "" || in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "email, username, and password are required")
	}

	// Check password length and requirements.
	if len(in.Password) < 8 || !domain.HasRequiredChars(in.Password) {
		return nil, status.Error(codes.InvalidArgument, "password must be at least 8 characters and contain uppercase, lowercase, number, and special character")
	}

	// Hash the password using bcrypt.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to hash password")
	}

	// Prepare the parameters for inserting a new user.
	params := data.InsertUserParams{
		Username:    sql.NullString{String: in.Username, Valid: true},
		Email:       sql.NullString{String: in.Email, Valid: true},
		EmailStatus: sql.NullBool{Bool: false, Valid: true}, // Default email status to false.
		PasHash:     sql.NullString{String: string(hashedPassword), Valid: true},
	}

	// Insert the new user into the database.
	user, err := s.Queries.InsertUser(ctx, params)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	// Return the ID of the created user in the response.
	return &ssov1.RegisterResponse{UserId: user.ID.String()}, nil
}
