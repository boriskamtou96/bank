package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secret string
}

func NewJWTMaker(secret string) (Maker, error) {
	if len(secret) < minSecretKeySize {
		return nil, fmt.Errorf("secret key is too short")
	}
	return &JWTMaker{
		secret: secret,
	}, nil
}

// CreateToken creates a new token for a specific username and duration
func (m *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload := &Payload{
		Username: username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(m.secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// VerifyToken checks if the token is valid or not and returns the username if valid
func (m *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(m.secret), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, fmt.Errorf("token has expired")) {
			return nil, fmt.Errorf("invalid token")
		}
		return nil, err
	}
	if !jwtToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return jwtToken.Claims.(*Payload), nil
}
