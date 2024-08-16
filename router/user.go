package router

import (
	"github.com/AsTao/meeting/api"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes() {

	RegistRoute(func(regPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		regPublic.POST("/login", userApi.Login)

		// rgAuthUser := rgAuth.Group("user")
		// rgAuthUser.GET("", func(ctx *gin.Context) {
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"data": []map[string]any{
		// 			{"id": 1, "name": "aa"},
		// 			{"id": 2, "name": "bb"},
		// 		},
		// 	})
		// })

		// rgAuthUser.GET("/:id", func(ctx *gin.Context) {
		// 	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		// 		"data": map[string]any{"id": 2, "name": "bb"},
		// 	})
		// })
	})
}
