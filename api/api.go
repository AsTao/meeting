package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary helloword
// @Schemes
// @Description do ping
// @Tags 测试
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /helloworld [get]
func Helloworld(g *gin.Context) {
	g.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg": "hello world",
	})
}
