package api

import "github.com/gin-gonic/gin"

var eng *gin.Engine

func init() {
	eng = gin.Default()

	eng.POST("/api/v1/passwords", GeneratePassword)
	eng.GET("/api/v1/passwords", GetPasswordList)
}

func Run() {
	if err := eng.Run("0.0.0.0:9000"); err != nil {
		panic(err)
	}
}
