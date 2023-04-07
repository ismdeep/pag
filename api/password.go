package api

import (
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
)

type Password struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

var passwords []Password

func init() {
	passwords = make([]Password, 0)
}

type GeneratePasswordRequest struct {
	Name string `json:"name"`
}

func GeneratePassword(c *gin.Context) {

	var req GeneratePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	for _, password := range passwords {
		if password.Name == req.Name {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "already exists",
			})
			return
		}
	}

	password := gofakeit.Password(true, true, true, false, false, 8)

	passwords = append(passwords, Password{
		Name:     req.Name,
		Password: password,
	})

	c.JSON(200, gin.H{})
}

func GetPasswordList(c *gin.Context) {
	c.JSON(200, passwords)
}
