package main

import (
	"github.com/hifat/sodium-api/internal/database"
	"github.com/hifat/sodium-api/internal/model/gormModel"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

func main() {
	db := database.PostgresDB()
	GormMigrate(db)
}

func GormMigrate(db *gorm.DB) {
	db.AutoMigrate(&gormModel.User{})
	db.AutoMigrate(&gormModel.RefreshToken{})
}
