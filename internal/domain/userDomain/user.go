package userDomain

import "github.com/google/uuid"

type UserRepository interface {
	GetFieldsByID(ID uuid.UUID, field string) (value interface{}, err error)
}
