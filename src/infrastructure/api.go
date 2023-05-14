package infrastructure

import (
	"fmt"

	"github.com/allen-utec/vota-api-users/src/application"
	"github.com/gin-gonic/gin"
)

func init() {
	application.UserService.Init(&UserRepository{})
}

func InitApi() {
	router := gin.Default()
	router.Use(corsMiddleware())

	router.GET("/users", GetAllUsersHandler)
	router.POST("/users", CreateUserHandler)

	router.Run(fmt.Sprintf(":%s", getEnv("PORT", "8081")))
}
