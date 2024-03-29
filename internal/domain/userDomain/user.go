package userDomain

import (
	"context"

	"github.com/google/uuid"
)

//go:generate mockgen -source=./user.go -destination=../../repository/userRepo/mockUser/mockUserRepo.go -package=mockUserRepo
type UserRepository interface {
	GetFieldsByID(ctx context.Context, ID uuid.UUID, field string) (value interface{}, err error)
	CheckExists(ctx context.Context, col, value string) (exists bool, err error)
}

type ResponseUser struct {
	ID       uuid.UUID `json:"ID"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Password string    `json:"-"`
}
