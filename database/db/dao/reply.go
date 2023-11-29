package dao

import (
	"VideoWeb2/database/db/model"
	"context"
	"gorm.io/gorm"
)

type ReplyDao struct {
	*gorm.DB
}

func NewReplyDao(ctx context.Context) *ReplyDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &ReplyDao{NewDBClient(ctx)}
}

func (dao *ReplyDao) CreateReply(reply *model.Reply) (err error) {
	err = DB.Model(&model.Reply{}).Create(&reply).Error
	return
}
