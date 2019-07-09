package models

type Collect struct {
	Model
	OpenId  string `json:"open_id"`
	TopicId int    `json:"topic_id"`
}

func AddCollect(collect Collect) error {
	if err := db.Create(&collect).Error; err != nil {
		return err
	}
	return nil
}

func DelCollect(topicId int, openId string) error {
	if err := db.Where("topic_id = ? AND open_id = ?", topicId, openId).Delete(&Collect{}).Error; err != nil {
		return err
	}
	return nil
}

func GetCollect(openId string) (*Collect, error) {
	var (
		collect Collect
		err     error
	)
	data := make(map[string]interface{})
	data["open_id"] = openId

	err = db.Where(data).First(&collect).Error
	if err != nil {
		return nil, err
	}
	return &collect, nil
}
