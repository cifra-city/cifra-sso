package authsrv

import (
	"context"

	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
)

// Logout - method for user logout, clears the JWT cookie.
func (s *AuthServer) Logout(ctx context.Context, in *ssov1.Empty) (*ssov1.Empty, error) {
	s.Log.Infof("user with %s success logout fron device", in.String())
	return &ssov1.Empty{}, nil
}
