package topic_service

type Topic struct {
	ExamId    string `json:"exam_id"`
	TopicName string `json:"topic_name"`
	OptionA   string `json:"option_a"`
	OptionB   string `json:"option_b"`
	OptionC   string `json:"option_c"`
	OptionD   string `json:"option_d"`
}
