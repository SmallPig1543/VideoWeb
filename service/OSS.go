package service

import (
	"VideoWeb2/conf"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"path/filepath"
)

func GetURL(key string) (string, error) {
	client, err := oss.New(conf.OSS_END_POINT, conf.OSS_ACCESS_KEY_ID, conf.OSS_ACCESS_KEY_SECRET)
	if err != nil {
		return "", errors.New("oss配置错误")
	}
	bucket, err := client.Bucket(conf.OSS_BUCKET)
	if err != nil {
		return "", errors.New("oss配置错误")
	}
	signedGetURL, _ := bucket.SignURL(key, oss.HTTPGet, 300)
	return signedGetURL, nil
}

func Upload(FileName string) (string, error) {
	client, err := oss.New(conf.OSS_END_POINT, conf.OSS_ACCESS_KEY_ID, conf.OSS_ACCESS_KEY_SECRET)
	if err != nil {
		return "", errors.New("oss配置错误")
	}
	// 获取存储空间
	bucket, err := client.Bucket(conf.OSS_BUCKET)
	if err != nil {
		return "", errors.New("oss配置错误")
	}
	// 获取扩展名
	ext := filepath.Ext(FileName)
	//将发送过来的文件路径转化为oss的存储路径
	objectName := "video/" + uuid.Must(uuid.NewRandom()).String() + ext
	err = bucket.PutObjectFromFile(objectName, FileName)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("文件有误")
	}
	return objectName, nil
}
