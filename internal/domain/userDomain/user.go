package userDomain

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetFieldsByID(ctx context.Context, ID uuid.UUID, field string) (value interface{}, err error)
}

type ResponseUser struct {
	ID       uuid.UUID `json:"ID"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Password string    `json:"-"`
}
