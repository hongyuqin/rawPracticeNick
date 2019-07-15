package topic_service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"rawPracticeNick/common"
	"rawPracticeNick/models"
	"rawPracticeNick/pkg/gredis"
	"rawPracticeNick/routers/api"
	"strconv"
)

func getBeginWrongTopic(req *api.TopicReq) (*Topic, error) {
	collects, err := models.GetCollects(req.AccessToken)
	if err != nil {
		logrus.Error("GetCollects error :", err)
		return nil, err
	}
	for _, collect := range collects {
		_, err = gredis.LPush(common.WRONG_TOPIC_LIST, strconv.Itoa(collect.TopicId))
		if err != nil {
			logrus.Error("lpush redis error :", err)
			return nil, err
		}
	}
	return getTopicByIndex(common.WRONG_TOPIC_LIST, req.AccessToken, 0)
}
func NextWrongTopic(req *api.TopicReq) (*Topic, error) {
	if req.IsBegin {
		return getBeginWrongTopic(req)
	}
	if req.Operate == common.OPERATE_LAST {
		return getTopicByIndex(common.COLLECT_LIST, req.AccessToken, req.CurrentIndex-1)
	}
	if req.Operate == common.OPERATE_NEXT {
		return getTopicByIndex(common.COLLECT_LIST, req.AccessToken, req.CurrentIndex+1)
	}
	return nil, errors.New("no topic")

}
