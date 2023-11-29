package dao

import (
	"VideoWeb2/database/db/model"
	"context"
	"gorm.io/gorm"
)

type VideoDao struct {
	*gorm.DB
}

func NewVideoDao(ctx context.Context) *VideoDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &VideoDao{NewDBClient(ctx)}
}

func (dao *VideoDao) CreateVideo(video *model.Video) (err error) {
	err = dao.DB.Model(&model.Video{}).Create(&video).Error
	return
}

func (dao *VideoDao) FindVideoByVid(vid uint) (video model.Video, err error) {
	err = dao.DB.Model(&model.Video{}).Where("id=?", vid).First(&video).Error
	return
}
