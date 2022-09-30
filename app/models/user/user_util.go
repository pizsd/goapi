package user

import (
	"github.com/gin-gonic/gin"
	"goapi/pkg/app"
	"goapi/pkg/database"
	"goapi/pkg/paginator"
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

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(User{}),
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return
}
