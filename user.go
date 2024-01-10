package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Admins is a struct that represents the admins table in the database
type Users struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

var jwtKey = []byte("my_secret_key")

func generateToken(user Users) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	})

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func insertUser(user Users) error {
	var results []Users

	err := supabaseClient.DB.From("Users").Insert(user).Execute(&results)

	if err != nil {
		return err
	}

	return nil
}

func registerUserHandler(c *gin.Context) {
	var newUser Users

	if error := c.ShouldBind(&newUser); error != nil {
		return
	}

	err := insertUser(newUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to insert user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User inserted successfully"})
}
