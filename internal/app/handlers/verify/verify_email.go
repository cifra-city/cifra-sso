package verify

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cifra-city/cifra-jwt"
	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/domain/entities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) VerifyEmail(ctx context.Context, in *ssov1.Empty) (*ssov1.Empty, error) {
	log := s.Log

	userID, err := cifra_jwt.VerificationJWT(ctx, log, s.Config.JWT.SecretKey)
	if err != nil {
		log.Error("Error getting user from JWT-token: %s", err)
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	user, err := s.Queries.GetUserByID(ctx, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.NotFound, "user not found")
	} else if err != nil {
		log.Errorf("error getting user: %v", err)
		return nil, status.Error(codes.Internal, "failed to retrieve user")
	}

	if user.EmailStatus {
		log.Errorf("user %s has already confirmed their email", user.Username)
		return nil, status.Error(codes.PermissionDenied, "user has already confirmed your email")
	}

	eve, err := s.ActionPermission.GetEvent(user.Username)
	if err != nil {
		log.Errorf("error checking in queue: %v", err)
		return nil, status.Error(codes.Internal, "failed to check in queue")
	}
	if eve != entities.ConfirmEmail {
		log.Errorf("user %s is not in the confirm email queue", user.Username)
		return nil, status.Error(codes.PermissionDenied, "user is not in the change confirm email queue")
	}

	stmt := data.UpdateEmailStatusByIDParams{
		ID:          userID,
		EmailStatus: true,
	}

	_, err = s.Queries.UpdateEmailStatusByID(ctx, stmt)
	if err != nil {
		log.Errorf("error updating email status for user %s: %v", user.Username, err)
		return nil, status.Error(codes.Internal, "Error updating email status for user")
	}

	log.Infof("email status updated for user: %s", user.Username)
	return &ssov1.Empty{}, nil
}
