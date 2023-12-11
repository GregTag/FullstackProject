package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var tokenLifetime = time.Duration(24) * time.Hour

type AuthManager struct {
	signingKey []byte
}

func NewAuthManager(signingKey []byte) *AuthManager {
	return &AuthManager{signingKey: signingKey}
}

func (am *AuthManager) MakeToken(user_id uint) (string, error) {
	expirationTime := time.Now().Add(tokenLifetime)
	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprint(user_id),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(am.signingKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (am *AuthManager) FetchToken(token string) (*map[string]string, error) {
	claims := jwt.RegisteredClaims{}
	jwt_token, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
		return am.signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !jwt_token.Valid {
		return nil, errors.New("invalid auth token")
	}

	return convertClaimsToMap(&claims), nil
}

func convertClaimsToMap(claims *jwt.RegisteredClaims) *map[string]string {
	sub := claims.Subject
	exp := claims.ExpiresAt
	result := map[string]string{
		"sub": sub,
		"exp": fmt.Sprint(exp),
	}
	return &result
}
