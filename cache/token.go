package cache

import "time"

// 把用户token保存在redis中
func SaveUserToken(token, userID string, exp time.Duration) error {
	err := RedisClient.Set(UserTokenKey(token), userID, exp).Err()
	return err
}

// 通过Token获取对应的userid
func GetUserByToken(token string) (string, error) {
	return RedisClient.Get(UserTokenKey(token)).Result()
}

// 删除用户登录token实现登出
func DelUserToken(token string) error {
	return RedisClient.Del(UserTokenKey(token)).Err()
}
