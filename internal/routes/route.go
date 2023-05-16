package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

type routes struct {
	orm    *gorm.DB
	router *gin.RouterGroup
}

func New(orm *gorm.DB, router *gin.RouterGroup) *routes {
	return &routes{
		orm,
		router,
	}
}

func (r routes) Register() {
	r.authRoute()
}
