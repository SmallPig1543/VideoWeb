package service

import (
	"VideoWeb2/database/db/dao"
	"VideoWeb2/database/db/model"
	"VideoWeb2/response"
	"VideoWeb2/types"
	"context"
	"time"
)

type CommentServ struct {
}

func (s *CommentServ) CommentCreate(ctx context.Context, req *types.CommentRequest) (interface{}, error) {
	u, err := types.GetUserInfo(ctx)
	if err != nil {
		return response.BadResponse("用户错误"), err
	}
	commentDao := dao.NewCommentDao(ctx)
	comment := model.Comment{
		Uid:      u.ID,
		Vid:      req.Vid,
		Content:  req.Content,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	err = commentDao.CreateComment(&comment)
	if err != nil {
		return response.BadResponse("存入数据库失败"), err
	}
	return response.SuccessResponse(), nil
}
