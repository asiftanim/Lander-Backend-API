package middlewares

import (
	"errors"
	"fmt"
	"lander/services"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("Authorization Header is not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("Invalid Authorization Header format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := errors.New("Authorization type not supported")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		accessToken := fields[1]
		token, err := services.ValidateToken(accessToken)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println("Claims[authorized]: ", claims["authorized"])
			fmt.Println("Claims[user_id]: ", claims["user_id"])
			fmt.Println("Claims[email]: ", claims["email"])
			fmt.Println("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		}

	}
}
