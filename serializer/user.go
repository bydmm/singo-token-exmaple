package serializer

import "singo/model"

// User 用户序列化器
type User struct {
	ID          uint   `json:"id"`
	UserName    string `json:"user_name"`
	Nickname    string `json:"nickname"`
	Status      string `json:"status"`
	Avatar      string `json:"avatar"`
	CreatedAt   int64  `json:"created_at"`
	Token       string `json:"token,omitempty"`
	TokenExpire int64  `json:"token_expire,omitempty"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
