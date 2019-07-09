package topic_service

import (
	"github.com/sirupsen/logrus"
	"rawPracticeNick/models"
)

func NextCollect(openId string) (*Topic, error) {
	collect, err := models.GetCollect(openId)
	if err != nil {
		logrus.Error("GetCollect error :", err)
		return nil, err
	}
	//要去topic表找具体题目
	topic, err := models.GetTopic(collect.TopicId)
	if err != nil {
		logrus.Error("GetTopic error :", collect.TopicId)
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
