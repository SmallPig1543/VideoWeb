package types

type CommentRequest struct {
	Vid     uint   `json:"vid" form:"vid"`
	Content string `json:"content" form:"content"`
}
