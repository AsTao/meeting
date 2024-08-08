package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes() {
	RegistRoute(func(regPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		regPublic.POST("/login", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "Login Success",
			})
		})
	})
}
