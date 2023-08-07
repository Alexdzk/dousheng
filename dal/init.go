package dal

import (
	"github.com/Alexdzk/dousheng/dal/cache"
	"github.com/Alexdzk/dousheng/dal/db"
	"github.com/Alexdzk/dousheng/dal/mq"
)

// Init init dal
func Init() {
	db.Init() //mysql
	cache.Init()
	mq.Init()
}
