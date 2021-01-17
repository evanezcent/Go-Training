package main

import (
	"fmt"

	"./configuration"
	"./controller"
	"./repository"
	"./service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = configuration.InitConnection()
	userRepository repository.UserRepository = repository.NewUserRepo(db)
	jwtService     service.JWTService        = service.NewJwtService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	userController controller.AuthController = controller.NewAuth(authService, jwtService)
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
