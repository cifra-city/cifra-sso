package authsrv

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cifra-city/cifra-sso/internal/db/data"
	"github.com/cifra-city/cifra-sso/internal/domain"
	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *AuthServer) CreateAdmin(ctx context.Context, in *ssov1.CreateAdminRequest) (*ssov1.CreateAdminResponse, error) {
	log := s.Log

	userID, err := s.Authenticate(ctx)
	if err != nil {
		return nil, err
	}
	CreatorID := domain.ToNullUUID(userID)
	if !CreatorID.Valid {
		log.Debugf("error while validating user %s", CreatorID.UUID.String())
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	note := domain.ToNullString(in.Name)
	username := domain.ToNullString(in.Username)

	newAdmin, err := s.Queries.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Debugf("user not foun %s", username.String)
			return nil, status.Error(codes.NotFound, "user not found")
		}
		log.Errorf("error while getting user %s", err.Error())
		return nil, status.Error(codes.Internal, "error getting user")
	}

	NewAdmin, err := s.Queries.InsertAdmin(ctx, data.InsertAdminParams{
		Uid:  domain.ToNullUUID(newAdmin.ID),
		Name: note,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("error while inserting user %s", err.Error())
		return nil, status.Error(codes.Internal, "Error insert data in a table")
	}
	if err == nil {
		log.Errorf("admin with this username: %s already exists", username.String)
		return nil, status.Error(codes.AlreadyExists, "Admin with this username already exists")
	}

	log.Infof("created new admin with id %s and username %s not about this admin: %s",
		newAdmin.ID, username.String, note.String)
	return &ssov1.CreateAdminResponse{Id: NewAdmin.ID}, nil
}
