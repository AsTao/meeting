package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 用户登录
// @Schemes
// @Description 用户登录详细描述
// @Tags 用户
// @Accept json
// @Produce json
// @Param  name formData string true "用户名"
// @Param  password formData string true "密码"
// @Success 200 {string} string	"ok" "登录成功"
// @Failure 400 {string} string "登录失败"
// @Router /user/login [post]
func Login(g *gin.Context) {
	g.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg": "Login Success",
	})
}
