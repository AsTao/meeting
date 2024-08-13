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

		rgAuthUser := rgAuth.Group("user")
		rgAuthUser.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"data": []map[string]any{
					{"id": 1, "name": "aa"},
					{"id": 2, "name": "bb"},
				},
			})
		})

		rgAuthUser.GET("/:id", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": map[string]any{"id": 2, "name": "bb"},
			})
		})
	})
}
