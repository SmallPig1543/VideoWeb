package es

import (
	"VideoWeb2/database/db/model"
	"VideoWeb2/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strconv"
)

// CreateMessageDocument 单个添加文档
func CreateMessageDocument(indexName string, message *model.Message) (err error) {
	DocID := strconv.Itoa(int(message.ID))
	_, err = EsClient.Index().Index(indexName).Id(DocID).BodyJson(*message).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// QueryMessageDocuments 查询时间范围内的消息
func QueryMessageDocuments(indexName string, filter *types.Filter) ([]model.Message, int64, error) {
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermQuery("PublisherID", strconv.Itoa(int(filter.Uid))))
	query = query.Must(elastic.NewRangeQuery("CreateAt").Gte(filter.StartTime).Lte(filter.EndTime))
	res, err := EsClient.Search(indexName).Query(query).From(0).Size(10).Do(context.Background())
	if err != nil {
		return nil, 0, err
	}
	count := res.Hits.TotalHits.Value // 获取返回条件个数
	var list []model.Message
	for _, v := range res.Hits.Hits {
		var obj model.Message
		err := json.Unmarshal(v.Source, &obj)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, obj)
	}
	return list, count, nil
}
