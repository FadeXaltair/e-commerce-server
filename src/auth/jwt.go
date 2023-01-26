package auth

import (
	"e-commerce-backend/config"
	"e-commerce-backend/database"
	"log"

	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// CustomClaim is a structure for the token
type CustomClaim struct {
	Name string
	Role string
	jwt.StandardClaims
}

// GenerateJwtToken function is used to create the token
func GenerateJwtToken(name string) (string, error) {
	mySignedKey := []byte("")

	User := &CustomClaim{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, User)
	ss, err := token.SignedString(mySignedKey)
	if err != nil {
		log.Println(err)
		return "", err

	}
	data := database.Tokens{
		Token: ss,
	}
	result := config.DB.Create(&data)
	if result.Error != nil {
		log.Println(err)
		return "", err
	}
	return ss, nil
}

// CheckJWTToken is used to validate the token
func CheckJWTToken(token string) (jwt.MapClaims, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	claims := jwtToken.Claims.(jwt.MapClaims)
	if err != nil {
		return claims, err
	}
	if jwtToken.Valid {
		return claims, nil
	}
	return claims, err
}

// Authorisation middleware
func Authorisaton(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "" {
		validtoken := strings.Split(token, " ")[1]
		claim, err := CheckJWTToken(validtoken)
		if err != nil {
			c.JSON(400, gin.H{
				"error":   true,
				"message": "error while checking",
			})
			return
		}
		c.Set("Name", claim["Name"])
		c.Next()
	} else {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "token required",
		})
	}
}
