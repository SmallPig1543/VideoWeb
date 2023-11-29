package es

import (
	"VideoWeb2/conf"
	"github.com/olivere/elastic/v7"
)

var EsClient *elastic.Client

func LinkEs() {
	client, err := elastic.NewClient(
		elastic.SetURL(conf.EsUrl),
		elastic.SetSniff(false), //是否开启集群嗅探
		elastic.SetBasicAuth("elastic", "123456"),
	)
	if err != nil {
		panic(err)
	}
	EsClient = client
}
