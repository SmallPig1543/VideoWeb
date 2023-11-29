package dao

import (
	"VideoWeb2/database/db/model"
	"context"
	"gorm.io/gorm"
)

type CommentDao struct {
	*gorm.DB
}

func NewCommentDao(ctx context.Context) *CommentDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &CommentDao{NewDBClient(ctx)}
}

func (dao *CommentDao) CreateComment(comment *model.Comment) (err error) {
	err = DB.Model(&model.Comment{}).Create(&comment).Error
	return
}
