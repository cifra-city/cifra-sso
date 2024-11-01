package services

import (
	"context"

	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
)

// Authentication defines methods for user authentication and authorization.
type Authentication interface {
	Login(ctx context.Context, in *ssov1.LoginReq) (*ssov1.LoginResp, error)
	Logout(ctx context.Context, in *ssov1.Empty) (*ssov1.Empty, error)
	ChangePassword(ctx context.Context, in *ssov1.ChangePassReq) (*ssov1.Empty, error)
	ChangeUsername(ctx context.Context, in *ssov1.ChangeUsernameReq) (*ssov1.Empty, error)
	ChangeEmail(ctx context.Context, in *ssov1.ChangeEmailReq) (*ssov1.Empty, error)
}
