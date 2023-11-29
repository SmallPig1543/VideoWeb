package dao

import (
	"VideoWeb2/database/db/model"
	"context"
	"gorm.io/gorm"
)

type MessageDao struct {
	*gorm.DB
}

func NewMessageDao(ctx context.Context) *MessageDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &MessageDao{NewDBClient(ctx)}
}

func (dao *MessageDao) CreateMessage(message *model.Message) (err error) {
	err = DB.Model(&model.Message{}).Create(&message).Error
	return
}
