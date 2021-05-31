package cache

import "fmt"

// UserTokenKey token的redis key
func UserTokenKey(token string) string {
	return fmt.Sprintf("user:token:%s", token)
}
