package model

type Video struct {
	ID       uint   `gorm:"primarykey" json:"ID"`
	Uid      uint   `json:"uid"` //发布者id
	Title    string `json:"title"`
	Types    string `json:"types"` //视频的种类
	Key      string //储存在oss中的key
	CreateAt string `json:"create_at"`
}

type VideoForEs struct {
	ID       uint   `json:"ID"`
	Uid      uint   `json:"uid"` //发布者id
	Title    string `json:"title"`
	Types    string `json:"types"` //视频的种类
	Views    int    `json:"views"`
	Key      string `json:"key"`
	CreateAt string `json:"CreateAt"`
}

func (Video) Mapping() string {
	return `{
  "mappings": {
    "properties": {
      "ID":{
        "type": "integer"
      },
      "uid":{
        "type": "integer"
      },
      "title": {
        "type": "text",
        "analyzer": "ik_smart"
      },
      "views":{
        "type":"integer"
      },
      "key":{
        "type": "keyword"
      },
      "types": {
        "type": "keyword"
      },
      "CreateAt": {
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}`
}
func (Video) IndexName() string {
	return "video"
}
