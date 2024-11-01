package services

import (
	"context"

	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
)

type Registration interface {
	Register(ctx context.Context, in *ssov1.RegisterReq) (*ssov1.Empty, error)
}
