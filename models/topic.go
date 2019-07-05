package models

import "time"

type Topic struct {
	Model
	TopicName      string    `json:"topic_name"`
	OptionA        string    `json:"option_a"`
	OptionB        string    `json:"option_b"`
	OptionC        string    `json:"option_c"`
	OptionD        string    `json:"option_d"`
	Answer         string    `json:"answer"`
	TopicAnalysis  string    `json:"topic_analysis"`
	WrongNum       int       `json:"wrong_num"`
	RightNum       int       `json:"right_num"`
	BelongTimeType time.Time `json:"belong_time_type"`
	ElementType    string    `json:"element_type"`
}
