package api

import (
	"singo/cache"
	"singo/serializer"
	"singo/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserTokenRefresh 用户刷新token接口
func UserTokenRefresh(c *gin.Context) {
	user := CurrentUser(c)
	var service service.UserTokenRefreshService
	res := service.Refresh(c, user)
	c.JSON(200, res)
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	// 移动端登出
	token := c.GetHeader("X-Token")
	if token != "" {
		_ = cache.DelUserToken(token)
	} else {
		// web端登出
		s := sessions.Default(c)
		s.Clear()
		s.Save()
	}
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
