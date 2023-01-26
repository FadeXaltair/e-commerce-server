package api

import (
	"e-commerce-backend/config"
	"e-commerce-backend/database"
	"e-commerce-backend/src/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// signup function is used to create the account in the database
func SignUp(c *gin.Context) {
	var userdata database.User
	err := c.Bind(&userdata)
	if err != nil {
		config.Error(err)
		return
	}
	confirmPass := userdata.ConfirmPassword
	password := userdata.Password
	hash, _ := auth.HashPassword(password)
	status := auth.CheckPasswordHash(confirmPass, hash)
	if !status {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "incorrect pass",
		})
		return
	}
	data := database.User{
		Name:     userdata.Name,
		Email:    userdata.Email,
		Password: hash,
	}
	err = CreateAccount(data)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error":   true,
			"message": "error while creating",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "account created successfully",
	})

}

// Login function is used to login the user
func Login(c *gin.Context) {
	var data database.Login
	err := c.Bind(&data)
	if err != nil {
		config.Error(err)
		return
	}
	var user database.User
	config.DB.Raw(`select x.id, x.name,x.password, x.email from products.users x where x.email =?`, data.Email).Scan(&user)
	if user.Id == 0 {
		config.Error(err)
		return
	}
	status := auth.CheckPasswordHash(data.Password, user.Password)
	if !status {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "incorrect password",
		})
		return
	}
	token, _ := auth.GenerateJwtToken(user.Name)
	resp := database.Response{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  resp,
		"token": token,
	})
}

func CreateAccount(data database.User) error {
	err := config.DB.Exec(`insert into products.users (user_name,email,password)
	values(?,?,?)`, data.Name, data.Email, data.Password).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
