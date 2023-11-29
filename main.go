package main

import (
	"VideoWeb2/database/cache"
	"VideoWeb2/database/db/dao"
	"VideoWeb2/database/db/model"
	"VideoWeb2/es"
	"VideoWeb2/mq"
	"VideoWeb2/route"
)

func init() {
	dao.InitMySQL()
	mq.LinkRabbitmq()
	cache.LinkRedis()
	es.LinkEs()
	es.CreateIndex(model.Message{}.IndexName(), model.Message{}.Mapping())
	es.CreateIndex(model.Video{}.IndexName(), model.Video{}.Mapping())
}
func main() {
	r := route.NewRouter()
	_ = r.Run(":9090")
}
