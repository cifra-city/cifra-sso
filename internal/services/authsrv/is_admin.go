package authsrv

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cifra-city/cifra-sso/internal/domain"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *AuthServer) IsAdmin(ctx context.Context, in *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	// Authenticate the user.
	userID, err := s.Authenticate(ctx)
	if err != nil {
		return nil, err
	}
	ID := domain.ToNullUUID(userID)
	if !ID.Valid {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}
	// Use the userID to check if the user is an admin.
	_, err = s.Queries.GetAdminByUID(ctx, ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &ssov1.IsAdminResponse{IsAdmin: false}, nil
		}

		return nil, status.Error(codes.Internal, "failed to verify admin status")
	}

	return &ssov1.IsAdminResponse{IsAdmin: true}, nil
}
