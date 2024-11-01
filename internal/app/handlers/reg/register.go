package reg

import (
	"context"
	"database/sql"
	"errors"

	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/pkg/security"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register - method for registering a new user.
func (s *Server) Register(ctx context.Context, in *ssov1.RegisterReq) (*ssov1.Empty, error) {
	log := s.Log

	log.Debugf("email: %s; user: %s; password: %s; ", in.Email, in.Username, in.Password)
	if in.Email == "" || in.Username == "" || in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "email, username, and password are required")
	}

	if len(in.Password) < 8 || !security.HasRequiredChars(in.Password) {
		return nil, status.Error(codes.InvalidArgument, "password must be at least 8 characters and contain uppercase, lowercase, number, and special character")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Debugf("error hashing password: %v", err)
		return nil, status.Error(codes.Internal, "failed to hash password")
	}

	_, err = s.Queries.GetUserByEmail(ctx, in.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("error getting user by email: %v", err)
		return nil, status.Error(codes.Internal, "Error getting user by email")
	}
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "This email address already exists")
	}

	_, err = s.Queries.GetUserByUsername(ctx, in.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("error getting user by username: %v", err)
		return nil, status.Error(codes.Internal, "Error getting user by username")
	}
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "This username already exists")
	}

	newUserID := uuid.New()

	// Prepare the parameters for inserting a new user.
	params := data.InsertUserParams{
		ID:          newUserID,
		Username:    in.Username,
		Email:       in.Email,
		EmailStatus: false,
		PassHash:    string(hashedPassword),
	}

	// Insert the new user into the database.
	user, err := s.Queries.InsertUser(ctx, params)
	if err != nil {
		log.Infof("error inserting user: %v", err)
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	log.Infof("user created: %v", user.Username)
	return &ssov1.Empty{}, nil
}
