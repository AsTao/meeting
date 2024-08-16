package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func NewUserApi() UserApi {
	return UserApi{}
}

// @Summary 用户登录
// @Description 用户登录详细描述
// @Param   name formData string true "用户名"
// @Param   password formData string true "密码"
// @Success 200 {string} string	"ok" "登录成功"
// @Failure 400 {string} string "登录失败"
func (m UserApi) Login(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg": "Login Success",
	})
}
