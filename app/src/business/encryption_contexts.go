package business

import "fmt"

func PasswordEncryptionContext(username string) string {
	return fmt.Sprintf("{username:%s,dataType:password}", username)
}

func RefreshTokenSecretEncryptionContext(username string) string {
	return fmt.Sprintf("{username:%s,dataType:refreshTokenSecret}", username)
}
