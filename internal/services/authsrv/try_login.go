package authsrv

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cifra-city/cifra-sso/internal/db/data"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *AuthServer) TryLogin(ctx context.Context, in *ssov1.TryLoginRequest) (*ssov1.TryLoginResponse, error) {
	log := s.Log
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
		code, err := s.Email.GenerateConfirmationCode()
		if err != nil {
			log.Errorf("Error gen code for list to email: %s", err)
			return nil, status.Error(codes.Internal, "error generating confirmation code")
		}
		s.Email.AddToEmailList(user.Username, code)

	}

	return &ssov1.TryLoginResponse{EmailStatus: user.EmailStatus}, nil
}
