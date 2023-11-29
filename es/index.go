package es

import (
	"context"
	"errors"
	"fmt"
)

//索引操作

// CreateIndex 创建索引 Dsl是创建索引的dsl语句
func CreateIndex(indexName string, Dsl string) (err error) {
	if exists := ExistsIndex(indexName); exists {
		return errors.New("索引已存在")
	}
	res, err := EsClient.CreateIndex(indexName).
		BodyString(Dsl).
		Do(context.Background())
	if err != nil {
		fmt.Println(res)
		return
	}
	return nil
}

// ExistsIndex 判断索引是否存在
func ExistsIndex(indexName string) bool {
	exists, _ := EsClient.IndexExists(indexName).Do(context.Background())
	return exists
}

// DeleteIndex 删除索引
func DeleteIndex(indexName string) (err error) {
	if exists := ExistsIndex(indexName); !exists {
		return errors.New("索引不存在")
	}
	res, err := EsClient.DeleteIndex(indexName).
		Do(context.Background())
	if err != nil {
		return
	}
	fmt.Println(res)
	return nil
}
