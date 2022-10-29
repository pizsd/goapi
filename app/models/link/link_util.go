package link

import (
	"github.com/gin-gonic/gin"
	"goapi/pkg/app"
	"goapi/pkg/cache"
	"goapi/pkg/database"
	"goapi/pkg/helpers"
	"goapi/pkg/paginator"
	"time"
)

func Get(idstr string) (link Link) {
	database.DB.Where("id", idstr).First(&link)
	return
}

func GetBy(field, value string) (link Link) {
	database.DB.Where(field+" = ?", field, value).First(&link)
	return
}

func All() (links []Link) {
	database.DB.Find(&links)
	return
}

func AllCached() (links []Link) {
	key := "links:all"
	expireTime := time.Second * 60
	cache.GetObject(key, &links)

	if helpers.Empty(links) {
		links = All()
		if helpers.Empty(links) {
			return
		}
		cache.Set(key, links, expireTime)
	}
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Link{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Link{}),
		&links,
		app.V1URL(database.TableName(&Link{})),
		perPage,
	)
	return
}
