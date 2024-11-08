package auth

import (
	"context"

	"github.com/cifra-city/cifra-jwt"
	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) LoginConfirm(ctx context.Context, in *ssov1.LoginConfirmReq) (*ssov1.LoginConfirmResp, error) {
	var user data.UsersSecret
	var err error

	log := s.Log

	if in.Username != nil {
		user, err = s.Queries.GetUserByUsername(ctx, *in.Username)
		if err != nil {
			log.Errorf("Error getting user by username in DB: %s", err)
			return nil, status.Error(codes.Internal, "failed to retrieve user")
		}
	}
	if in.Email != nil {
		user, err = s.Queries.GetUserByEmail(ctx, *in.Email)
		if err != nil {
			log.Errorf("Error getting user by rmail in DB: %s", err)
			return nil, status.Error(codes.Internal, "failed to retrieve user")
		}
	}

	availability := s.Email.CheckInEmailList(user.Email, in.Code)
	if !availability {
		log.Errorf("Code is not available")
		return nil, status.Error(codes.PermissionDenied, "code is not available")
	}

	token, err := cifra_jwt.GenerateJWT(user.ID, s.Config.JWT.TokenLifetime, s.Config.JWT.SecretKey)
	if err != nil {
		log.Error("error creating jwt token %s", err)
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	log.Infof("successfully logged in user %s", user.Username)
	return &ssov1.LoginConfirmResp{
		Token: &token,
	}, nil
}
