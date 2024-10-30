package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/modules/jwt"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Login - method for user login and JWT token generation.
func (s *AuthServer) Login(ctx context.Context, in *ssov1.LoginReq) (*ssov1.LoginResp, error) {
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

	//TODO ADD CHECK FOR USER EMAIL STATUS AND ADD VER BY EMAIL

	// Compare the provided password with the stored hashed password.
	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(in.Password))
	if err != nil {
		log.Debugf("incorect password for user: %s", user.Username)
		return nil, status.Error(codes.Unauthenticated, "invalid password")
	}

	// Generate a JWT token.
	token, err := jwt.GenerateJWT(user.ID, s.Config.JWT.TokenLifetime, s.Config.JWT.SecretKey)
	if err != nil {
		log.Error("error creating jwt token %s", err)
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	// Return the token in the response The frontend has to process it
	log.Infof("successfully logged in user %s", user.Username)
	return &ssov1.LoginResp{Token: token}, nil
}
