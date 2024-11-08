package verify

import (
	"context"

	"github.com/cifra-city/cifra-jwt"
	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
	"github.com/cifra-city/cifra-sso/internal/app/modules/email"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SendConfirmCode is a method that sends a confirmation email to the user
func (s *Server) SendConfirmCode(ctx context.Context, in *ssov1.Empty) (*ssov1.InquiryResp, error) {
	log := s.Log

	user, err := cifra_jwt.VerificationJWT(ctx, log, s.Config.JWT.SecretKey)
	if err != nil {
		log.Errorf("Error getting user from JWT-token: %s", err)
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	userDB, err := s.Queries.GetUserByID(ctx, user)
	if err != nil {
		log.Errorf("Error getting user by ID in DB: %s", err)
		return nil, status.Error(codes.Internal, "failed to retrieve user")
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

	return &ssov1.InquiryResp{Link: "TODO"}, nil
}
