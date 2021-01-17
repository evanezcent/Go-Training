package controller

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthController interface for login register user
type AuthController interface{
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct{

}

// NewAuth is like constructor of the model
func NewAuth() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "Logged In",
	})
}

func (c *authController) Register(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "Registered",
	})
}