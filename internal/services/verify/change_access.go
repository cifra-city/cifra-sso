package verify

import (
	"context"

	"github.com/cifra-city/cifra-sso/internal/modules/jwt"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
)

// AccessForChanges is a method that adds a user to the event list
func (s *VerServer) AccessForChanges(ctx context.Context, in *ssov1.AccessReq) (*ssov1.AccessResp, error) {
	log := s.Log

	user, err := jwt.VerificationJWT(ctx, log, s.Config.JWT.SecretKey)
	if err != nil {
		log.Error("Error getting user from JWT-token: %s", err)
		return nil, err
	}

	userDB, err := s.Queries.GetUserByID(ctx, user)
	if err != nil {
		log.Error("Error getting user by ID in DB: %s", err)
		return nil, err
	}

	if !s.Email.CheckInEmailList(userDB.Username, in.Code) {
		log.Infof("Users code is incorect: %s", err)
		return nil, err
	}

	s.Events.AddToQueue(userDB.Username, in.Eve.String())

	log.Infof("Add user %s to event list %s", userDB.Username, in.Eve.String())

	return &ssov1.AccessResp{}, nil
}
