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

// CreateVideoDocument 单个添加文档
func CreateVideoDocument(indexName string, video *model.VideoForEs) (err error) {
	DocID := strconv.Itoa(int(video.ID))

	//jsonStr := fmt.Sprintf(`{"CreateAt":"%s","ID":%d,"title":"%s","types":"%s","uid":%d,"views":%d}`, video.CreateAt, video.ID, video.Title, video.Types, video.Uid, 0)
	//str := "`" + jsonStr + "`"
	//fmt.Println(str)
	//_, err = EsClient.Index().Index(indexName).Id(DocID).BodyJson(str).Do(context.Background())
	//_, err = EsClient.Index().Index(indexName).Id(DocID).BodyString(str).Do(context.Background())
	_, err = EsClient.Index().Index(indexName).Id(DocID).BodyJson(*video).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// QueryVideoDocuments 查询符合条件的video
func QueryVideoDocuments(indexName string, filter *types.Filter) ([]model.VideoForEs, int64, error) {
	query := elastic.NewBoolQuery()
	if filter.Vid != 0 {
		fmt.Println(filter.Vid)
		query = query.Must(elastic.NewTermQuery("ID", filter.Vid))
	}
	if filter.Uid != 0 {
		fmt.Println(filter.Uid)
		query = query.Must(elastic.NewTermQuery("uid", filter.Uid))
	}
	if filter.Types != "" {
		fmt.Println(filter.Types)
		query = query.Must(elastic.NewMatchQuery("types", filter.Types))
	}
	if filter.StartTime != "" && filter.EndTime != "" {
		query = query.Must(elastic.NewRangeQuery("CreateAt").Gte(filter.StartTime).Lte(filter.EndTime))
	}
	if filter.Views != 0 {
		query = query.Must(elastic.NewRangeQuery("views").Gte(filter.Views))
	}
	res, err := EsClient.Search(indexName).Query(query).From(0).Size(10).Do(context.Background())
	if err != nil {
		return nil, 0, err
	}
	count := res.Hits.TotalHits.Value // 获取返回条件个数
	var list []model.VideoForEs
	for _, v := range res.Hits.Hits {
		var obj model.VideoForEs
		err := json.Unmarshal(v.Source, &obj)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, obj)
	}
	return list, count, nil
}

func UpdateViews(DocID string, value int) error {
	res, err := EsClient.Update().Index(model.Video{}.IndexName()).Id(DocID).Doc(map[string]any{
		"views": value,
	}).Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("%#v", res)
	return nil
}
