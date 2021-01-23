package main

import (
	"fmt"

	"./configuration"
	"./controller"
	"./middleware"
	"./repository"
	"./service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = configuration.InitConnection()
	userRepository repository.UserRepository = repository.NewUserRepo(db)
	bookRepository repository.BookRepository = repository.NewBookRepo(db)
	jwtService     service.JWTService        = service.NewJwtService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	userService    service.UserService       = service.NewUserService(userRepository)
	bookService    service.BookService       = service.NewBookService(bookRepository)
	userController controller.AuthController = controller.NewAuth(authService, jwtService, userService)
	bookController controller.BookController = controller.NewBookController(bookService, jwtService)
)

func main() {
	fmt.Println("Starting apps...")

	r := gin.Default()

	authRoutes := r.Group("api/user")
	{
		authRoutes.POST("/login", userController.Login)
		authRoutes.POST("/register", userController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/get", userController.Get)
		userRoutes.PUT("/update", userController.Update)
	}

	bookRoutes := r.Group("api/books", middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("/", bookController.All)
		bookRoutes.GET("/book/:id", bookController.FindByID)
		bookRoutes.POST("/add", bookController.Insert)
		bookRoutes.PUT("/update/:id", bookController.Update)
		bookRoutes.DELETE("/delete/:id", bookController.Delete)
	}

	r.Run()
}
