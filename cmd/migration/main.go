package main

import (
	"github.com/hifat/hifat-blog-api/internal/database"
	"github.com/hifat/hifat-blog-api/internal/model/gormModel"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

func main() {
	db := database.PostgresDB()
	GormMigrate(db)
}

func GormMigrate(db *gorm.DB) {
	db.AutoMigrate(&gormModel.User{})
}
