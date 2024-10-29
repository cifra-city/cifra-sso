package auth

import (
	"context"

	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// InquiryForChange is a method that sends a confirmation email to the user
func (s *AuthServer) InquiryForChange(ctx context.Context, in *ssov1.InquiryReq) (*ssov1.InquiryResp, error) {
	log := s.Log

	user, err := s.Authenticate(ctx)
	if err != nil {
		log.Error("Error getting user from JWT-token: %s", err)
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	userDB, err := s.Queries.GetUserByID(ctx, user)
	if err != nil {
		log.Error("Error getting user by ID in DB: %s", err)
		return nil, err
	}
	code, err := s.Email.GenerateConfirmationCode()
	if err != nil {
		log.Error("Error generating confirmation code: %s", err)
		return nil, err
	}

	err = s.Email.SendConfirmationEmail(userDB.Email, code)
	if err != nil {
		log.Error("Error sending confirmation email: %s", err)
		return nil, err
	}

	return &ssov1.InquiryResp{Link: "TODO"}, nil
}
