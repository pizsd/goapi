package factories

import (
	"github.com/bxcodec/faker/v3"
	"goapi/app/models/link"
)

func MakeLinks(count int) []link.Link {

	var objs []link.Link

	// 设置唯一性，如 Link 模型的某个字段需要唯一，即可取消注释
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		linkModel := link.Link{
			Name: faker.Word(),
			Url:  faker.URL(),
			Logo: faker.URL(),
		}
		objs = append(objs, linkModel)
	}

	return objs
}
