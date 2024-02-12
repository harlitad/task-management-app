package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"example.com/task-management-app/model"
	"example.com/task-management-app/pkg/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header is required",
			})
			ctx.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			fmt.Println(authParts)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization header",
			})
			ctx.Abort()
			return
		}

		token, claims, err := auth.VerifyToken(authParts[1])

		if err != nil || token == false {
			fmt.Println(err)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Set("user", model.UserAuthenticated{
			Email: claims["email"].(string),
			Id:    claims["id"].(string),
			Name:  claims["name"].(string),
		})

		ctx.Set("userId", claims["id"])

		ctx.Next()
	}

}
