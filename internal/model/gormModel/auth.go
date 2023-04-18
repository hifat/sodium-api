package gormModel

import (
	"time"

	"github.com/google/uuid"
	"github.com/hifat/sodium-api/internal/utils/gorm/utype"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID       uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()" json:"ID"`
	Token    string    `gorm:"type:text;unique" json:"token"`
	Agent    string    `gorm:"type:varchar(100)" json:"agent"`
	ClientIP utype.IP  `gorm:"type:text" json:"clientIP"`
	IsActive bool      `gorm:"boolean; default:true" json:"isActive"`

	UserID uuid.UUID `gorm:"type:uuid" json:"userID"`
	User   User      `json:"user"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
