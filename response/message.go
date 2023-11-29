package response

type MessageResp struct {
	PublisherID uint   `json:"publisher_id"`
	ReceiverID  uint   `json:"receiver_id"`
	Content     string `json:"content"`
	CreateAt    string `json:"create_at"`
}
