package models

import "time"

type Topic struct {
	Model
	TopicName      string `json:"topic_name"`
	OptionA        string `json:"option_a"`
	OptionB        string `json:"option_b"`
	OptionC        string `json:"option_c"`
	OptionD        string `json:"option_d"`
	Answer         string `json:"answer"`
	TopicAnalysis  string `json:"topic_analysis"`
	WrongNum       int    `json:"wrong_num"`
	RightNum       int    `json:"right_num"`
	Region         string `json:"region"`
	ExamType       string `json:"exam_type"`
	ElementTypeOne string `json:"element_type_one"`
	ElementTypeTwo string `json:"element_type_two"`
}

func AddTopic(topic *Topic) error {
	topic.CreateTime = time.Now()
	topic.UpdateTime = time.Now()
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
	if topic.Region != "" {
		data["region"] = topic.Region
	}
	if topic.ElementTypeOne != "" {
		data["element_type_one"] = topic.ElementTypeOne
	}
	if topic.ElementTypeTwo != "" {
		data["element_type_two"] = topic.ElementTypeTwo
	}

	err = db.Where(data).Find(&topics).Error
	if err != nil {
		return nil, err
	}
	return topics, nil
}
