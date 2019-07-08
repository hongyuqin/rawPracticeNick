package models

type WrongTopic struct {
	Model
	OpenId  string `json:"open_id"`
	TopicId int    `json:"topic_id"`
}

func AddWrongTopic(wrongTopic WrongTopic) error {
	if err := db.Create(&wrongTopic).Error; err != nil {
		return err
	}
	return nil
}

func DelWrongTopic(id int) error {
	data := make(map[string]interface{})
	data["id"] = id
	if err := db.Model(&WrongTopic{}).Delete(data).Error; err != nil {
		return err
	}
	return nil
}
