package main

import (
	"e-commerce-backend/config"
	"e-commerce-backend/database"
	"e-commerce-backend/src/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// configuring database connection
	config.DB = database.DatabaseConnection()
	// dbconn is used to close the conection after the use of database
	dbConn, err := config.DB.DB()
	if err != nil {
		log.Println(err)
	}
	//Closing the database connection
	defer dbConn.Close()
	// defining our router// The url pointing to API definition
	r := gin.Default()

	// here wer run all the routes
	routes.Routes(r)
}
