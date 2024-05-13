package middlewares

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.Abort()

			utils.Response(c, &utils.ApiUnauthorized, nil, "access token is missing")
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			log.Println(err.Error())

			c.Abort()

			utils.Response(c, &utils.ApiUnauthorized, nil, "access token is not valid")
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if float64(time.Now().UnixMilli()) > claims["exp"].(float64) {
			c.Abort()

			utils.Response(c, &utils.ApiUnauthorized, nil, "access token is expired")
			return
		}

		c.Set("id", claims["id"])
		c.Set("email", claims["email"])

		c.Next()
	}
}
