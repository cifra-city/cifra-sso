package authsrv

import (
	"context"

	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
)

func (s *AuthServer) AccessForChanges(ctx context.Context, in *ssov1.AccessReq) (*ssov1.AccessResp, error) {
	log := s.Log

	user, err := s.Authenticate(ctx)
	if err != nil {
		log.Error("Error getting user from JWT-token: %s", err)
		return nil, err
	}

	userDB, err := s.Queries.GetUserByID(ctx, user)
	if err != nil {
		log.Error("Error getting user by ID in DB: %s", err)
		return nil, err
	}

	if s.Email.CheckInEmailList(userDB.Username, in.Code) {
		log.Infof("Users code is incorect: %s", err)
		return nil, err
	}

	return &ssov1.AccessResp{
		Link: "link TODO",
	}, nil
}
