package types

type MessageReq struct {
	ReceiverID uint   `json:"receiver_id" form:"receiver_id"`
	Content    string `json:"content" form:"content"`
}
