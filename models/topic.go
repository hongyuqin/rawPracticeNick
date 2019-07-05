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

func AddTopic(topic Topic) error {
	if err := db.Create(&topic).Error; err != nil {
		return err
	}
	return nil
}

func GetTopic(id int) (*Topic, error) {
	var topic Topic
	if err := db.Where("id = ? AND flag = 0", id).First(&topic).Error; err != nil {
		return nil, err
	}
	return &topic, nil
}

func GetTopics(topic *Topic) ([]Topic, error) {
	var (
		topics []Topic
		err    error
	)
	data := make(map[string]interface{})
	if topic.ElementType != "" {
		data["element_type"] = topic.ElementType
	}
	if !topic.BelongTimeType.IsZero() {
		data["belong_time_type"] = topic.BelongTimeType
	}
	err = db.Where(topic).Find(topics).Error
	if err != nil {
		return nil, err
	}
	return topics, nil
}
func UpdateTopic(topic *Topic) error {
	data := make(map[string]interface{})
	if topic.ElementType != "" {
		data["element_type"] = topic.ElementType
	}
	if !topic.BelongTimeType.IsZero() {
		data["belong_time_type"] = topic.BelongTimeType
	}
	if err := db.Model(&Topic{}).Where("id = ? AND flag = 0 ", topic.ID).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
