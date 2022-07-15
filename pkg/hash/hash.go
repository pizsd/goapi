package hash

import (
	"goapi/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func BcryptHash(pwd string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	logger.LogIf(err)
	return string(bytes)
}

func BcryptCheck(pwd, hashedPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	return err == nil
}

func BcryptIsHashed(str string) bool {
	return len(str) == 60
}
