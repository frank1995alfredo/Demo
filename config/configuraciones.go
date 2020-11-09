package config

import (
	"net/http"

	//_ "github.com/dgrijalva/jwt-go"

	_ "github.com/gin-contrib/cors" //gsdg

	//metodos "github.com/frank1995alfredo/api/controllers/mantenimiento/metodos"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/*********************CORS*****************************************/

// CORS Middleware
func CORS(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

//HashPassword ... encripta el password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash ... hace un check del password y el hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
