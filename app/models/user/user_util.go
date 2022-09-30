package user

import (
	"goapi/pkg/database"
)

func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(&User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(&User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

func GetByMulti(loginId string) (userModel User) {
	database.DB.Where("phone = ?", loginId).
		Or("email = ?", loginId).
		Or("name = ?", loginId).First(&userModel)
	return
}

func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

func Find(idStr string) (userModel User) {
	database.DB.Where("id", idStr).First(&userModel)
	return
}

func All() (users []User) {
	database.DB.Find(&users)
	return users
}
