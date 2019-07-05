package models

type User struct {
	Model
	UserName       string `json:"username"`
	PassWord       string `json:"password"`
	OpenId         string `json:"open_id"`
	HasLearnNum    int    `json:"has_learn_num"`
	DailyNeedNum   int    `json:"daily_need_num"`
	AnswerDuration int    `json:"answer_duration"`
	PracticeDays   int    `json:"practice_days"`
	PracticeTime   int    `json:"practice_time"`
	AnswerNum      int    `json:"answer_num"`
}

func AddUser(user User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
func SelectUserByOpenId(openId string) (*User, error) {
	var user User
	err := db.Where("open_id = ? AND flag = 0", openId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
