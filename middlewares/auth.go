package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
// func AuthorizeJWT() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		const BEARER_SCHEMA = "Bearer "
// 		authHeader := c.GetHeader("Authorization")
// 		tokenString := authHeader[len(BEARER_SCHEMA):]

// 		token, err := ValidateToken(tokenString)

// 		if token.Valid {
// 			claims := token.Claims.(jwt.MapClaims)
// 			log.Println("Claims[Name]: ", claims["authorized"])
// 			log.Println("Claims[Admin]: ", claims["user_id"])
// 			log.Println("Claims[ExpiresAt]: ", claims["exp"])
// 		} else {
// 			log.Println(err)
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		}
// 	}
// }

// func ValidateToken(tokenString string) (*jwt.Token, error) {
// 	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Signing method validation
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		// Return the secret signing key
// 		return []byte(jwtSrv.secretKey), nil
// 	})
// }

