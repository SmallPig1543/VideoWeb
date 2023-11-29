package types

import (
	"context"
	"errors"
)

type UserRequest struct {
	UserName string `json:"user_name" form:"user_name" binding:"required,max=100"`
	PassWord string `json:"pass_word" form:"pass_word" binding:"required,min=3,max=100"`
}

// UserUploadRequest 上传或者修改头像URL
type UserUploadRequest struct {
	AvatarURL string `form:"avatarURL"`
}

var userKey int

type UserInfo struct {
	ID uint `json:"id"`
}

// NewContext 新创建一个context，存有userInfo
func NewContext(c context.Context, user *UserInfo) context.Context {
	return context.WithValue(c, userKey, user)
}

// GetUserInfo 从context中获取userInfo
func GetUserInfo(c context.Context) (*UserInfo, error) {
	user, ok := c.Value(userKey).(*UserInfo)
	if !ok {
		return nil, errors.New("获取用户信息失败")
	}
	return user, nil
}
