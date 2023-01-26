package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DB is a type of gorm db
var DB *gorm.DB

// database credentials
var (
	DBHost     = "localhost"
	DBPort     = "5432"
	DBUser     = "bogi"
	DBPassword = "root"
	DBName     = "stq_dev"
	DBSSL      = "disable"
)

// Error handling
func Error(err error) gin.H {
	log.Println(err)
	return gin.H{
		"statuscode": 400,
		"error":      true,
		"message":    "error occured",
	}
}
