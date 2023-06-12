package repository

import (
	"github.com/google/wire"
	"github.com/hifat/sodium-api/internal/database"
	"gorm.io/gorm"
)

var GormDBSet = wire.NewSet(NewGormDB)

func NewGormDB() (*gorm.DB, func()) {
	orm := database.PostgresDB()
	db, err := orm.DB()
	if err != nil {
		panic(err)
	}

	close := func() {
		db.Close()
	}

	return orm, close
}
