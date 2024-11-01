package services

import (
	"context"

	ssov1 "github.com/cifra-city/cifra-sso/internal/api"
)

type Verification interface {
	InquiryForChange(ctx context.Context, in *ssov1.InquiryReq) (*ssov1.InquiryResp, error)
	AccessForChanges(ctx context.Context, in *ssov1.AccessReq) (*ssov1.AccessResp, error)
}
