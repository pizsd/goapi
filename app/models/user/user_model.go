package user

import (
	"goapi/app/models"
	"goapi/pkg/database"
	"goapi/pkg/hash"
)

type User struct {
	models.BaseModel
	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`
	models.CommonTimestampsField
}

func (u *User) Create() {
	database.DB.Create(&u)
}

func (u *User) ComparePassword(pwd string) bool {
	return hash.BcryptCheck(pwd, u.Password)
}

func (u *User) Save() (rowsAffected int64) {
	res := database.DB.Save(&u)
	return res.RowsAffected
}
