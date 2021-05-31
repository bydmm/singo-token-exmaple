package service

import (
	"singo/model"
	"singo/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Token    bool   `form:"token" json:"token"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if !user.CheckPassword(service.Password) {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	var token string
	var tokenExpire int64
	var err error
	if service.Token {
		token, tokenExpire, err = user.MakeToken()
		if err != nil {
			return serializer.DBErr("redis err", err)
		}
	} else {
		// web端设置session
		service.setSession(c, user)
	}

	data := serializer.BuildUser(user)
	data.Token = token
	data.TokenExpire = tokenExpire
	return serializer.Response{
		Data: data,
	}
}
