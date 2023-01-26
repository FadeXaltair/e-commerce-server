package database

import (
	"e-commerce-backend/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// postgresql database connection
func DatabaseConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		//database host
		config.DBHost,
		//database port
		config.DBPort,
		//database user
		config.DBUser,
		//database pass
		config.DBPassword,
		//database name
		config.DBName,
		//ssl
		config.DBSSL)), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	return db
}
