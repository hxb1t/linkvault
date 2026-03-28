package domain

import "fmt"

func UserSessionKey(username string) string {
	return fmt.Sprintf("session:user:%s", username)
}
