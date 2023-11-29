package dao

import (
	"VideoWeb2/database/db/model"
	"context"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = DB.Model(&model.User{}).Create(&user).Error
	return
}

func (dao *UserDao) FindUserByUserName(username string) (user *model.User, err error) {
	err = DB.Model(&model.User{}).Where("user_name=?", username).First(&user).Error
	return
}
func (dao *UserDao) FindUserByUid(uid uint) (user *model.User, err error) {
	err = DB.Model(&model.User{}).Where("id=?", uid).First(&user).Error
	return
}

func (dao *UserDao) UploadAvatar(URL string, uid uint) (err error) {
	user, _ := dao.FindUserByUid(uid)
	user.AvatarURL = URL
	err = dao.DB.Save(user).Error
	return
}
