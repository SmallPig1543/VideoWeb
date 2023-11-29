package model

// Message 聊天消息的结构体
type Message struct {
	ID          uint `gorm:"primarykey"`
	PublisherID uint
	ReceiverID  uint
	Content     string
	CreateAt    string
}

func (Message) IndexName() string {
	return "message"
}

// Mapping 生成索引
func (Message) Mapping() string {
	return `{
  "mappings": {
    "properties": {
      "ID":{
        "type": "integer"
      },
      "content": {
        "type": "text",
        "analyzer": "ik_smart"
      },
      "PublisherID": {
        "type": "integer"
      },
      "ReceiverID": {
        "type": "integer"
      },
      "CreateAt": {
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}
