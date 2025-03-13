// auth/auth.go
package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JwtWrapper wraps the signing key and issuer information
type JwtWrapper struct {
	SecretKey         string
	Issuer            string
	AccessExpiration  int64
	RefreshExpiration int64
}

// JwtClaims adds username and user ID to the standard JWT claims
type JwtClaims struct {
	Username string `json:"username"`
	UserId   uint   `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateTokens generates both access and refresh tokens
func (j *JwtWrapper) GenerateTokens(username string, userId uint) (string, string, error) {
	// Generate Access Token
	accessClaims := &JwtClaims{
		Username: username,
		UserId:   userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(j.AccessExpiration))),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccessToken, err := accessToken.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", "", err
	}

	// Generate Refresh Token
	refreshClaims := &JwtClaims{
		UserId:   userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(j.RefreshExpiration))),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", "", err
	}

	return signedAccessToken, signedRefreshToken, nil
}

// VerifyToken verifies the access token and extracts claims
func (j *JwtWrapper) VerifyToken(signedToken string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// Ensure the token's signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// VerifyRefreshToken verifies the refresh token and checks if it is valid for generating a new access token
func VerifyRefreshToken(userId uint, refreshToken string) (bool, error) {
	jwtWrapper := JwtWrapper{
		SecretKey: os.Getenv("JWT_SECRET"), // Use environment variable for the secret
	}

	claims, err := jwtWrapper.VerifyToken(refreshToken)
	if err != nil || claims.UserId != userId {
		return false, errors.New("invalid or expired refresh token")
	}

	return true, nil
}

// VerifyRefreshTokenByUsername verifies the refresh token and checks if it is valid for generating a new access token
func VerifyRefreshTokenByUsername(username string, refreshToken string) (bool, error) {
	jwtWrapper := JwtWrapper{
		SecretKey: os.Getenv("JWT_SECRET"), // Use environment variable for the secret
	}

	// Verify the token and extract claims
	claims, err := jwtWrapper.VerifyToken(refreshToken)
	if err != nil || claims.Username != username {
		return false, errors.New("invalid or expired refresh token")
	}

	return true, nil
}
