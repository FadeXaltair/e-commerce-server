package routes

import (
	"e-commerce-backend/src/api"
	"e-commerce-backend/src/auth"

	"github.com/gin-gonic/gin"
)

// Routes function is used for the routing for our apis
func Routes(router *gin.Engine) {
	router.POST("/signup", api.SignUp)
	router.POST("/login", api.Login)
	router.GET("/products", api.Products)
	auth := router.Group("/secured").Use(auth.Authorisaton)
	{
		auth.POST("/cart", api.AddToCart)
		auth.POST("/order-purchased", api.OrderPlaced)
		auth.GET("/orders", api.Orders)
		auth.GET("/cart", api.CartData)
	}
	router.Run()
}
