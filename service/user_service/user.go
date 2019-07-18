package user_service

import (
	"github.com/sirupsen/logrus"
	"rawPracticeNick/common"
	"rawPracticeNick/models"
	"rawPracticeNick/pkg/gredis"
	"strconv"
)

/**
		"rank":10,//排名
     	"total":100,//总人数
    	"left_days":,//考试剩余天数
    	"daily_need_num":,//每日需刷题数量
    	"today_practice_num":,//今日刷题数量 存redis
    	"has_learn_num":,//已学数量
    	"wrong_num":10,//做错数量 需复习数量
*/
type HomeDetail struct {
	Rank             int    `json:"rank"`
	Total            int    `json:"total"`
	LeftDays         int    `json:"left_days"`
	TodayPracticeNum int    `json:"today_practice_num"`
	DailyNeedNum     int    `json:"daily_need_num"`
	HasLearnNum      int    `json:"has_learn_num"`
	WrongNum         int    `json:"wrong_num"`
	TotalTopicNum    int    `json:"total_topic_num"`
	Region           string `json:"region"`
}

func Home(openId string) (*HomeDetail, error) {
	//0.获取用户基本信息
	user, err := models.SelectUserByOpenId(openId)
	if err != nil {
		logrus.Error("select user error :", err)
		return nil, err
	}
	//TODO 1.获取排名
	//2.获取总人数 TOTAL
	count, err := models.CountUser(user.Region, user.ExamType)
	if err != nil {
		logrus.Error("count user error :", err)
		return nil, err
	}
	//TODO:3.获取剩余天数
	//4.获取今日做题数
	todayNumStr, err := gredis.Get(common.TODAY_PREFIX + openId)
	if err != nil {
		logrus.Error("redis Get error :", err)
		todayNumStr = []byte("0")
	}
	todayNum, err := strconv.Atoi(string(todayNumStr))
	if err != nil {
		logrus.Error("Atoi error :", err)
		return nil, err
	}
	//5.今日错题数
	topics, err := models.GetWrongTopics(openId)
	if err != nil {
		logrus.Error("GetWrongTopics error :", err)
		return nil, err
	}
	homeDetail := &HomeDetail{
		Rank:             1,
		Total:            count,
		TodayPracticeNum: todayNum,
		DailyNeedNum:     user.DailyNeedNum,
		HasLearnNum:      user.HasLearnNum,
		WrongNum:         len(topics),
		Region:           user.Region,
		TotalTopicNum:    1000,
	}
	return homeDetail, nil
}

type PlanReq struct {
	AccessToken  string `schema:"accessToken"`
	Region       string `schema:"region"`
	ExamType     string `schema:"exam_type"`
	DailyNeedNum int    `schema:"daily_need_num"`
}

func Plan(req *PlanReq) error {
	if err := models.UpdateUser(&models.User{
		OpenId:       req.AccessToken,
		Region:       req.Region,
		ExamType:     req.ExamType,
		DailyNeedNum: req.DailyNeedNum,
	}); err != nil {
		logrus.Error("updateUser error :", err)
		return err
	}
	return nil
}
