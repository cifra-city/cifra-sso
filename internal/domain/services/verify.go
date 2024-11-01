package services

import (
	"context"

	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
)

type Verification interface {
	SendConfirmCode(ctx context.Context, in *ssov1.Empty) (*ssov1.InquiryResp, error)
	AccessForChanges(ctx context.Context, in *ssov1.AccessReq) (*ssov1.AccessResp, error)
}
