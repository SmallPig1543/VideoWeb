package types

type ReplyRequest struct {
	Cid     uint   `json:"cid" form:"cid"`
	Content string `json:"content" form:"content"`
}
