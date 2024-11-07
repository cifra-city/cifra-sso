package verify

import (
	"context"

	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
	"github.com/cifra-city/cifra-sso/internal/pkg/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CheckConfirmCode is a method that adds a user to the event list
func (s *Server) CheckConfirmCode(ctx context.Context, in *ssov1.CheckConfirmCodeReq) (*ssov1.CheckConfirmCodeResp, error) {
	log := s.Log

	user, err := jwt.VerificationJWT(ctx, log, s.Config.JWT.SecretKey)
	if err != nil {
		log.Error("Error getting user from JWT-token: %s", err)
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	userDB, err := s.Queries.GetUserByID(ctx, user)
	if err != nil {
		log.Error("Error getting user by ID in DB: %s", err)
		return nil, status.Error(codes.Internal, "failed to retrieve user")
	}

	if !s.Email.CheckInEmailList(userDB.Email, in.Code) {
		log.Infof("User`s %s code is incorect: %s", userDB.Email, in.Code)
		return nil, status.Error(codes.PermissionDenied, "failed to check in email list")
	}

	s.ActionPermission.AddToQueue(userDB.Username, in.Eve.String())

	log.Infof("Add user %s to event list %s", userDB.Username, in.Eve.String())

	return &ssov1.CheckConfirmCodeResp{
		Eve: in.Eve,
	}, nil
}
