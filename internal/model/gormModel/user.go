package gormModel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Username     string    `gorm:"type:varchar(100);unique" json:"username"`
	Password     string    `gorm:"type:varchar(100);" json:"password"`
	Name         string    `gorm:"type:varchar(100);" json:"name"`
	RefreshToken string    `gorm:"type:varchar(100);" json:"refreshToken"`
	gorm.Model
}
