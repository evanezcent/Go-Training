package controller

import (
	"net/http"
	"strconv"

	"../dto"
	"../helper"
	"../model"
	"../service"
	"github.com/gin-gonic/gin"
)

// AuthController interface for login register user
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

// NewAuth is like constructor of the model
func NewAuth(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService,
		jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	err := ctx.ShouldBind(&loginDTO)

	if err != nil {
		res := helper.ResponseFailed("Failed to process", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)

		return
	}

	authRes := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if val, ok := authRes.(model.User); ok {
		generateToken := c.jwtService.GenerateToken(strconv.FormatUint(val.ID, 10))
		val.Token = generateToken

		res := helper.ResponseSucces(true, "success", val)

		ctx.JSON(http.StatusOK, res)

		return
	}

	response := helper.ResponseFailed("Invalid Credential", "Invalid Credential", nil)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var newUser dto.UserCreateDTO
	err := ctx.ShouldBind(&newUser)
	if err != nil {
		res := helper.ResponseFailed("Failed to process", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	if c.authService.IsDuplicateEmail(newUser.Email) {
		res := helper.ResponseFailed("Email has ben registered", "Failed", nil)
		ctx.JSON(http.StatusConflict, res)
	} else {
		createUser := c.authService.CreateUser(newUser)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createUser.ID, 10))
		createUser.Token = token

		res := helper.ResponseSucces(true, "success", createUser)
		ctx.JSON(http.StatusOK, res)
	}
}
