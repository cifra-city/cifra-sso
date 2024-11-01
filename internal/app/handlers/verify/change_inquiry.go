package verify

import (
	"context"

	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
	"github.com/cifra-city/cifra-sso/internal/app/modules/email"
	"github.com/cifra-city/cifra-sso/internal/pkg/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// InquiryForChange is a method that sends a confirmation email to the user
func (s *Server) InquiryForChange(ctx context.Context, in *ssov1.InquiryReq) (*ssov1.InquiryResp, error) {
	log := s.Log

	user, err := jwt.VerificationJWT(ctx, log, s.Config.JWT.SecretKey)
	if err != nil {
		log.Error("Error getting user from JWT-token: %s", err)
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	userDB, err := s.Queries.GetUserByID(ctx, user)
	if err != nil {
		log.Error("Error getting user by ID in DB: %s", err)
		return nil, err
	}
	code, err := email.GenerateConfirmationCode()
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
