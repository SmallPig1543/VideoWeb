package service

import (
	"VideoWeb2/database/db/dao"
	"VideoWeb2/database/db/model"
	"VideoWeb2/response"
	"VideoWeb2/types"
	"VideoWeb2/util"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
}

func (s *UserService) Register(ctx context.Context, req *types.UserRequest) (interface{}, error) {
	userDao := dao.NewUserDao(ctx)
	u, err := userDao.FindUserByUserName(req.UserName)
	if err == nil {
		return response.BadResponse("该用户已存在"), errors.New("该用户已存在")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u = &model.User{
			UserName: req.UserName,
			PassWord: req.PassWord,
			CreateAt: time.Now().Format("2006-01-02 15:03:04"),
		}
		if err = u.SetPassword(req.PassWord); err != nil {
			return response.BadResponse("密码不符要求"), err
		}
		if err = userDao.CreateUser(u); err != nil {
			return response.BadResponse("存入数据库失败"), err
		}
	}

	return response.SuccessResponse(), nil
}

func (s *UserService) Login(ctx context.Context, req *types.UserRequest) (interface{}, error) {
	userDao := dao.NewUserDao(ctx)
	u, err := userDao.FindUserByUserName(req.UserName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.BadResponse("该用户不存在"), err
	}
	if !u.CheckPassword(req.PassWord) {
		err = errors.New("密码错误")
		return response.BadResponse("密码错误"), err
	}

	token, err := util.GenerateToken(u.ID, u.UserName)
	if err != nil {
		return response.BadResponse("token签发失败"), err
	}
	userResp := &response.TokenData{
		User: &response.UserResp{
			ID:       u.ID,
			UserName: u.UserName,
			CreateAt: u.CreateAt,
		},
		Token: token,
	}
	return response.SuccessResponseWithData(userResp), nil
}

func (s *UserService) Upload(ctx context.Context, req *types.UserUploadRequest) (interface{}, error) {
	userDao := dao.NewUserDao(ctx)
	u, err := types.GetUserInfo(ctx)
	if err != nil {
		return response.BadResponse("token有误"), err
	}
	err = userDao.UploadAvatar(req.AvatarURL, u.ID)
	if err != nil {
		return response.BadResponse("上传/修改失败"), err
	}
	return response.SuccessResponse(), nil
}
