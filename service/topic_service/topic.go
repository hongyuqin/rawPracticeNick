package topic_service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"rawPracticeNick/common"
	"rawPracticeNick/models"
	"rawPracticeNick/pkg/gredis"
	"rawPracticeNick/pkg/util"
)

type Topic struct {
	TopicId   int    `json:"topic_id"`
	TopicName string `json:"topic_name"`
	OptionA   string `json:"option_a"`
	OptionB   string `json:"option_b"`
	OptionC   string `json:"option_c"`
	OptionD   string `json:"option_d"`
}

//提交答案的返回
type AnswerResp struct {
	MyAnswer      string `json:"my_answer"`
	Right         bool   `json:"right"`
	Answer        string `json:"answer"`
	TopicAnalysis string `json:"topic_analysis"`
	WrongNum      int    `json:"wrong_num"`
	RightNum      int    `json:"right_num"`
}

func NextTopic(openId string) (*Topic, error) {
	//-1.TODO 拿到用户当前设置
	allTopicsMap := make(map[int]struct{})

	region := ""
	elementTypeOne := ""
	elementTypeTwo := ""
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
		allTopicsMap[topic.ID] = struct{}{}
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
		delete(allTopicsMap, topic.ID)
		//TODO:存到redis
	}
	//3.把这两个值存到redis 未做的题目：已做的题目 随机抽一题出来

	//4.过滤一下 就是要做的题目表
	//biubiu: 因为放到map达到了随机的效果
	for id := range allTopicsMap {
		topic, err := models.GetTopic(id)
		if err != nil {
			logrus.Error("GetTopic error ", id)
			return nil, err
		}
		//5.返回题目
		return &Topic{
			TopicId:   topic.ID,
			TopicName: topic.TopicName,
			OptionA:   topic.OptionA,
			OptionB:   topic.OptionB,
			OptionC:   topic.OptionC,
			OptionD:   topic.OptionD,
		}, nil
	}
	return nil, errors.New("no topic")

}
func Answer(openId string, topicId int, myAnswer string) (*AnswerResp, error) {
	//1.看下答案对不对
	topic, err := models.GetTopic(topicId)
	if err != nil {
		logrus.Error("GetTopic error :", err)
		return nil, err
	}
	if topic == nil {
		logrus.Error("GetTopic empty :", topicId)
		return nil, err
	}
	//2.对的话更新题目的正确数，错的话更新题目错误数
	right := true
	if myAnswer == topic.Answer {
		topic.RightNum = topic.RightNum + 1
		//往用户错题表删除数据
		if err = models.DelWrongTopic(models.WrongTopic{
			OpenId:  openId,
			TopicId: topicId,
		}); err != nil {
			logrus.Error("DelWrongTopic error :", err)
			return nil, err
		}
	} else {
		right = false
		topic.WrongNum = topic.WrongNum + 1
		//往用户错题表插入数据
		if err = models.AddWrongTopic(models.WrongTopic{
			OpenId:  openId,
			TopicId: topicId,
		}); err != nil {
			logrus.Error("AddWrongTopic error :", err)
			return nil, err
		}
	}
	if err = models.UpdateTopic(topic); err != nil {
		logrus.Error("UpdateTopic error :", err)
		return nil, err
	}
	//3.往用户做过题表插入数据
	if err = models.AddDoneTopic(models.DoneTopic{
		OpenId:  openId,
		TopicId: topicId,
	}); err != nil {
		logrus.Error("AddDoneTopic error :", err)
		return nil, err
	}
	//3-1 存一个今日已答题
	if err = gredis.SAdd(common.TODAY_FINISH_PREFIX+openId, string(topicId)); err != nil {
		logrus.Error("AddDoneTopic sadd error ", err)
		return nil, err
	}
	//再设一个过期时间
	if err = gredis.ExpireAt(common.TODAY_FINISH_PREFIX+openId, util.GetNextDayBegin()); err != nil {
		logrus.Error("AddDoneTopic expireAt error ", err)
		return nil, err
	}

	//4.返回
	return &AnswerResp{
		MyAnswer:      myAnswer,
		Right:         right,
		Answer:        topic.Answer,
		TopicAnalysis: topic.TopicAnalysis,
		WrongNum:      topic.WrongNum,
		RightNum:      topic.RightNum,
	}, nil

}
func Collect(openId string, topicId int) error {
	if err := models.AddCollect(models.Collect{
		OpenId:  openId,
		TopicId: topicId,
	}); err != nil {
		logrus.Error("Collect error :", err)
		return err
	}
	return nil
}
