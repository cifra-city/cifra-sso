package auth

import (
	"context"

	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
	"github.com/cifra-city/cifra-sso/internal/app/modules/email"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/domain/entities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SendLoginCode(ctx context.Context, in *ssov1.SendLoginCodeReq) (*ssov1.Empty, error) {
	var userDB data.UsersSecret
	var err error

	log := s.Log

	if in.Username == nil {
		userDB, err = s.Queries.GetUserByUsername(ctx, *in.Username)
		if err != nil {
			log.Errorf("Error getting user by username in DB: %s", err)
			return nil, status.Error(codes.Internal, "failed to retrieve user")
		}
	}
	if in.Email == nil {
		userDB, err = s.Queries.GetUserByEmail(ctx, *in.Email)
		if err != nil {
			log.Errorf("Error getting user by rmail in DB: %s", err)
			return nil, status.Error(codes.Internal, "failed to retrieve user")
		}
	}

	code, err := email.GenerateConfirmationCode()
	if err != nil {
		log.Errorf("Error generating confirmation code: %s", err)
		return nil, status.Error(codes.Internal, "failed to generate confirmation code")
	}
	log.Debugf("Generated confirmation code: %s, %s", code, userDB.Email)

	err = s.Email.SendConfirmationCode(userDB.Email, code)
	if err != nil {
		log.Errorf("Error sending confirmation email: %s", err)
		return nil, status.Error(codes.Internal, "failed to send confirmation email")
	}

	s.ActionPermission.AddToQueue(userDB.Username, entities.Login)

	return &ssov1.Empty{}, nil
}
