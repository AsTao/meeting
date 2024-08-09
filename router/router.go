package router

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type IFnRegistRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRoutes []IFnRegistRoute
)

func RegistRoute(fn IFnRegistRoute) {
	if fn == nil {
		return
	}

	gfnRoutes = append(gfnRoutes, fn)
}

func InitBasePlatformRoutes() {
	InitUserRoutes()
}

func InitRouter() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:6666"}, // 允许的来源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Colg!")
	})

	rgPublic := r.Group("api/v1/public")
	rgAuth := r.Group("api/v1")

	InitBasePlatformRoutes()

	for _, fnRegistRoute := range gfnRoutes {
		fnRegistRoute(rgPublic, rgAuth)
	}
	port := viper.GetString("server.port")
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("Error starting server: %s", err)
	}

}
