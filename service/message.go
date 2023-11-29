package service

import (
	"VideoWeb2/database/db/dao"
	"VideoWeb2/database/db/model"
	"VideoWeb2/es"
	"VideoWeb2/mq"
	"VideoWeb2/response"
	"VideoWeb2/types"
	"context"
	"time"
)

type ChatServ struct {
}

func (s *ChatServ) SendMessage(ctx context.Context, req *types.MessageReq) (interface{}, error) {
	publisher, err := types.GetUserInfo(ctx)
	userDao := dao.NewUserDao(ctx)
	messageDao := dao.NewMessageDao(ctx)
	if err != nil {
		return response.BadResponse("用户出错"), err
	}
	receiver, err := userDao.FindUserByUid(req.ReceiverID)
	if err != nil {
		return response.BadResponse("不存在接收方"), err
	}
	err = mq.CreatePublisher(receiver.UserName, req.Content)
	if err != nil {
		return response.BadResponse("消息发送失败"), err
	}
	messages := make(chan string)
	channel := make(chan bool)

	go mq.Receive(receiver.UserName, messages, channel)
	<-channel
	for v := range messages {
		message := &model.Message{
			ID:          0,
			PublisherID: publisher.ID,
			ReceiverID:  req.ReceiverID,
			Content:     v,
			CreateAt:    time.Now().Format("2006-01-02 15:04:05"),
		}

		err = messageDao.CreateMessage(message)
		if err != nil {
			return response.BadResponse("数据库存入失败"), err
		}
		err = es.CreateMessageDocument(model.Message{}.IndexName(), message)
		if err != nil {
			return response.BadResponse("es创建文档失败"), err
		}
	}
	return response.SuccessResponse(), nil
}

func (s *ChatServ) QueryMessages(ctx context.Context, filter *types.Filter) (interface{}, error) {
	u, _ := types.GetUserInfo(ctx)
	filter.Uid = u.ID
	messages, count, err := es.QueryMessageDocuments(model.Message{}.IndexName(), filter)
	if err != nil {
		return response.BadResponse("es查询失败"), err
	}
	list := make([]*response.MessageResp, 0)
	for _, v := range messages {
		list = append(list, &response.MessageResp{
			PublisherID: v.PublisherID,
			ReceiverID:  v.ReceiverID,
			Content:     v.Content,
			CreateAt:    v.CreateAt,
		})
	}
	return response.ListResp(list, count), nil
}
