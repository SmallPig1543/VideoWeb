package service

import (
	"VideoWeb2/database/db/dao"
	"VideoWeb2/database/db/model"
	"VideoWeb2/response"
	"VideoWeb2/types"
	"context"
	"time"
)

type ReplyServ struct {
}

func (s *ReplyServ) ReplyCreate(ctx context.Context, req *types.ReplyRequest) (interface{}, error) {
	u, err := types.GetUserInfo(ctx)
	if err != nil {
		return response.BadResponse("用户错误"), err
	}
	replyDao := dao.NewReplyDao(ctx)
	reply := model.Reply{
		Uid:      u.ID,
		Cid:      req.Cid,
		Content:  req.Content,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	err = replyDao.CreateReply(&reply)
	if err != nil {
		return response.BadResponse("存入数据库失败"), err
	}
	return response.SuccessResponse(), nil
}
