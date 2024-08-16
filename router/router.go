package router

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	docs "github.com/AsTao/meeting/docs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	rgPublic := r.Group("api/v1/public")
	rgAuth := r.Group("api/v1")

	InitBasePlatformRoutes()

	for _, fnRegistRoute := range gfnRoutes {
		fnRegistRoute(rgPublic, rgAuth)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	port := viper.GetString("server.port")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("start server error: %s \n", err.Error())
		}
	}()

	<-ctx.Done()
	stop()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown: ", err)
		return
	}

	if err := r.Run(":" + port); err != nil {
		fmt.Printf("Error starting server: %s", err)
	}

}
