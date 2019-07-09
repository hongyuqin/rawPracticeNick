package topic_service

import (
	"github.com/sirupsen/logrus"
	"rawPracticeNick/models"
)

func NextWrongTopic(openId string) (*Topic, error) {
	wrongTopic, err := models.GetWrongTopic(openId)
	if err != nil {
		logrus.Error("GetTopics error :", err)
		return nil, err
	}
	//要去topic表找具体题目
	topic, err := models.GetTopic(wrongTopic.TopicId)
	if err != nil {
		logrus.Error("GetTopic error :", wrongTopic.TopicId)
		return nil, err
	}
	return &Topic{
		TopicId:   topic.ID,
		TopicName: topic.TopicName,
		OptionA:   topic.OptionA,
		OptionB:   topic.OptionB,
		OptionC:   topic.OptionC,
		OptionD:   topic.OptionD,
	}, nil

}
