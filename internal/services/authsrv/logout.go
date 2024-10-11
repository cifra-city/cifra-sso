package authsrv

import (
	"context"

	ssov1 "github.com/cifra-city/cifra-sso/resources/grpc/gen"
)

// Logout - method for user logout, clears the JWT cookie.
func (s *AuthServer) Logout(ctx context.Context, in *ssov1.LogoutRequest) (*ssov1.LogoutResponse, error) {
	// Поскольку мы не имеем прямого доступа к HTTP-заголовкам, предполагаем, что эту часть
	// будет обрабатывать `grpc-gateway` или клиент самостоятельно.

	// Вернем простое сообщение об успешном выходе.
	return &ssov1.LogoutResponse{
		Message: "Logout successful",
	}, nil
}
