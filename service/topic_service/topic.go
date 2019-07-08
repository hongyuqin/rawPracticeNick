package topic_service

import (
	"github.com/sirupsen/logrus"
	"rawPracticeNick/models"
)

type Topic struct {
	ExamId    string `json:"exam_id"`
	TopicName string `json:"topic_name"`
	OptionA   string `json:"option_a"`
	OptionB   string `json:"option_b"`
	OptionC   string `json:"option_c"`
	OptionD   string `json:"option_d"`
}

func BeginAnswer(openId, region, elementTypeOne, elementTypeTwo string) (*Topic, error) {
	//0.生成examId
	//1.拿到分类的全部id
	topics, err := models.GetTopics(&models.Topic{
		Region:         region,
		ElementTypeOne: elementTypeOne,
		ElementTypeTwo: elementTypeTwo,
	})
	if err != nil {
		logrus.Error("GetTopics error :", err)
		return nil, err
	}
	topicIdCol := make([]int, 0)
	for _, topic := range topics {
		topicIdCol = append(topicIdCol, topic.ID)
		//TODO: 存到redis
	}
	//2.拿到已做题目的全部id
	doneTopics, err := models.GetDoneTopics(openId)
	if err != nil {
		logrus.Error("GetDoneTopics error :", err)
		return nil, err
	}
	doneTopicIdCol := make([]int, 0)
	for _, topic := range doneTopics {
		doneTopicIdCol = append(doneTopicIdCol, topic.ID)
		//TODO:存到redis
	}
	//3.把这两个值存到redis 未做的题目：已做的题目 随机抽一题出来
	topic := topics[0]
	//4.过滤一下 就是要做的题目表

	//5.返回题目
	return &Topic{
		TopicName: topic.TopicName,
		OptionA:   topic.OptionA,
		OptionB:   topic.OptionB,
		OptionC:   topic.OptionC,
		OptionD:   topic.OptionD,
	}, nil
}
