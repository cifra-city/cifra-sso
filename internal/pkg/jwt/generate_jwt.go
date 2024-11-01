package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// GenerateJWT creates a JWT token for the authenticated user.
func GenerateJWT(userID uuid.UUID, tlt time.Duration, sk string) (string, error) {
	// Define token expiration time.
	expirationTime := time.Now().Add(tlt * time.Second)

	// Create the JWT claims, which include the user ID and expiration time.
	claims := &jwt.RegisteredClaims{
		Subject:   userID.String(), // Use user ID as the subject in string format.
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	// Create the JWT token using the claims and the secret key.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(sk))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
