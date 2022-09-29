package factories

import (
	"github.com/bxcodec/faker/v3"
	"goapi/app/models/user"
)

func MakeUsers(times int) []user.User {
	var objs []user.User
	faker.SetGenerateUniqueValues(true)
	for i := 0; i < times; i++ {
		obj := user.User{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Phone:    faker.Phonenumber(),
			Password: "$2a$14$wF0veezH5Hloe7YSPmAQ0uJCBhj.H1y43M.heXsCtHXBYyx/rUt8q",
		}
		objs = append(objs, obj)
	}
	return objs
}
