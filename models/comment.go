package models

type Comment struct {
	Model
	TopicId        int    `json:"topic_id"`
	CustomerId     int    `json:"customer_id"`
	CommentContent string `json:"comment_content"`
}
