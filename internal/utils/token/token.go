package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type JWT interface {
	CreateToken(userPayload UserPayload, duration time.Duration) (token string, payload *Payload, err error)
	VerifyToken(token string) (payload *Payload, err error)
}

type UserPayload struct {
	UserID   uuid.UUID `json:"userID"`
	Username string    `json:"username"`
}

type Payload struct {
	ID        uuid.UUID `json:"ID"`
	UserID    uuid.UUID `json:"userID"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiredAt time.Time `josn:"expireAt"`
}

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

func NewPayload(userPayload UserPayload, duration time.Duration) (payload *Payload, err error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload = &Payload{
		ID:        tokenID,
		UserID:    userPayload.UserID,
		Username:  userPayload.Username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
