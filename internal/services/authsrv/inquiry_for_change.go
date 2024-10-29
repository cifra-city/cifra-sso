package authsrv

import (
	"context"

	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
)

func (s *AuthServer) InquiryForChange(ctx context.Context, in *ssov1.InquiryReq) (*ssov1.InquiryResp, error) {
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

	s.Email.AddToEmailList(userDB.Username, in.Eve.String())
	log.Infof("Add usernamr to event queue %s, %s", userDB.Username, in.Eve.String())
	return &ssov1.InquiryResp{
		Eve: in.Eve,
	}, nil
}
