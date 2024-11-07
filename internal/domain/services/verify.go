package services

import (
	"context"

	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
)

type Verification interface {
	SendConfirmCode(ctx context.Context, in *ssov1.Empty) (*ssov1.InquiryResp, error)
	CheckConfirmCode(ctx context.Context, in *ssov1.CheckConfirmCodeReq) (*ssov1.CheckConfirmCodeResp, error)
	VerifyEmail(ctx context.Context, in *ssov1.Empty) (*ssov1.Empty, error)
}
