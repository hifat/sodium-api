package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	minSecretKeySize = 32
)

var ErrSecretKeySize = fmt.Errorf("invalid key size: must be at least %d charecters", minSecretKeySize)

type JWTToken struct {
	secretKey string
}

func NewJWTToken(secretKey string) (JWT, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, ErrSecretKeySize
	}

	return &JWTToken{secretKey}, nil
}

func (t JWTToken) CreateToken(userPayload UserPayload, duration time.Duration) (token string, payload *Payload, err error) {
	payload, err = NewPayload(userPayload, duration)
	if err != nil {
		return "", nil, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err = jwtToken.SignedString([]byte(t.secretKey))

	return token, payload, err
}

func (t JWTToken) VerifyToken(token string) (payload *Payload, err error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(t.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}

		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
