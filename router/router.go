package router

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/AsTao/meeting/api"
	"github.com/AsTao/meeting/docs"
	"github.com/AsTao/meeting/global"
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

func InitRouter() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		v1.GET("/helloworld", api.Helloworld)
		v1.POST("/user/login", api.Login)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	port := viper.GetString("server.port")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	go func() {
		global.Logger.Info(fmt.Sprintf(" Start Listen %s", port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("start server error: %s \n", err.Error()))
		}
	}()

	<-ctx.Done()
	stop()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("Server forced to shutdown: %s", err.Error()))
		return
	}

	if err := r.Run(":" + port); err != nil {
		global.Logger.Error(fmt.Sprintf("Error starting server: %s", err.Error()))
	}

}
