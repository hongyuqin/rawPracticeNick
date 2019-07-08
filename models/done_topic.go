package models

type DoneTopic struct {
	Model
	OpenId  string `json:"open_id"`
	TopicId int    `json:"topic_id"`
}

func AddDoneTopic(doneTopic DoneTopic) error {
	if err := db.Create(&doneTopic).Error; err != nil {
		return err
	}
	return nil
}

func DelDoneTopic(id int) error {
	data := make(map[string]interface{})
	data["id"] = id
	if err := db.Model(&DoneTopic{}).Delete(data).Error; err != nil {
		return err
	}
	return nil
}

func GetDoneTopics(openId string) ([]DoneTopic, error) {
	var (
		topics []DoneTopic
		err    error
	)
	data := make(map[string]interface{})
	data["open_id"] = openId

	err = db.Where(data).Find(&topics).Error
	if err != nil {
		return nil, err
	}
	return topics, nil
}
