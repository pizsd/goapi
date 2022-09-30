package factories

import (
	"github.com/bxcodec/faker/v3"
	"goapi/app/models/category"
	"goapi/pkg/console"
)

func MakeCategories(count int) []category.Category {

	var objs []category.Category

	// 设置唯一性，如 Category 模型的某个字段需要唯一，即可取消注释
	faker.SetGenerateUniqueValues(true)
	err := faker.SetRandomStringLength(4)
	console.ExitIf(err)
	for i := 0; i < count; i++ {
		categoryModel := category.Category{
			Name:        faker.ChineseName(),
			Description: faker.Sentence(),
		}
		objs = append(objs, categoryModel)
	}

	return objs
}
