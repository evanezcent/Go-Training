package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"./configuration"
	"./controller"
)

var (
	db *gorm.DB = configuration.InitConnection()
	userController controller.AuthController = controller.NewAuth()
)

func main() {
	fmt.Println("Starting apps...")

	r := gin.Default()
	
	authRoutes := r.Group("api/user")
	{
		authRoutes.POST("/login", userController.Login)
		authRoutes.POST("/register", userController.Register)
	}

	r.Run()
}
