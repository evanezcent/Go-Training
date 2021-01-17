package middleware

import (
	"log"
	"net/http"

	"../service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthorizeJWT to authorize jwt in api
func AuthorizeJWT(jwtServ service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			res := helper.responseFailed("Failed to process request", "Token invalid", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		token, err := jwtServ.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Printf("Claims[userID]: ", claims["userID"])
			log.Printf("Claims[issuer]: ", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.responseFailed("Token invalid", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
