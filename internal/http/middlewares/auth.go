package middlewares

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

// AuthMiddleware is a middleware to check if the request is authorized
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")

		log.Info().Msgf("Token: %s", token)

		tokenSplit := strings.Split(token, " ")

		if len(tokenSplit) != 2 {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if tokenSplit[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		parseToken, err := jwt.Parse(tokenSplit[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if !parseToken.Valid {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if claims, ok := parseToken.Claims.(jwt.MapClaims); !ok {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		} else {
			c.Set("user_id", claims["sub"])
		}

		c.Next()
	}
}
