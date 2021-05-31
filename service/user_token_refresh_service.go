package service

import (
	"singo/model"
	"singo/serializer"

	"github.com/gin-gonic/gin"
)

// UserTokenRefreshService 用户刷新token的服务
type UserTokenRefreshService struct {
}

// Refresh 刷新token
func (service *UserTokenRefreshService) Refresh(c *gin.Context, user *model.User) serializer.Response {
	token, tokenExpire, err := user.MakeToken()
	if err != nil {
		return serializer.DBErr("redis err", err)
	}
	data := serializer.BuildUser(*user)
	data.Token = token
	data.TokenExpire = tokenExpire
	return serializer.Response{
		Data: data,
	}
}
