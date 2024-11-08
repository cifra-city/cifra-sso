package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cifra-city/cifra-jwt"
	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
	"github.com/cifra-city/cifra-sso/internal/app/modules/email"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Login - method for user login and JWT token generation.
func (s *Server) Login(ctx context.Context, in *ssov1.LoginReq) (*ssov1.LoginResp, error) {
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

	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(in.Password))
	if err != nil {
		log.Debugf("incorect password for user: %s", user.Username)
		return nil, status.Error(codes.InvalidArgument, "invalid password")
	}

	if user.EmailStatus {
		code, err := email.GenerateConfirmationCode()
		if err != nil {
			log.Errorf("Error generating confirmation code: %s", err)
			return nil, status.Error(codes.Internal, "failed to generate confirmation code")
		}

		err = s.Email.SendConfirmationCode(user.Email, code)
		if err != nil {
			log.Errorf("Error sending confirmation email: %s", err)
			return nil, status.Error(codes.Internal, "failed to send confirmation email")
		}

		log.Debugf("Generated confirmation code: %s, %s", code, user.Email)
		return &ssov1.LoginResp{}, nil
	}

	token, err := cifra_jwt.GenerateJWT(user.ID, s.Config.JWT.TokenLifetime, s.Config.JWT.SecretKey)
	if err != nil {
		log.Error("error creating jwt token %s", err)
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	log.Infof("successfully logged in user %s", user.Username)
	return &ssov1.LoginResp{Token: &token}, nil
}
