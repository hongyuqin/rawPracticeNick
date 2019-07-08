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

func DelCollect(id int) error {
	data := make(map[string]interface{})
	data["id"] = id
	if err := db.Delete(data).Error; err != nil {
		return err
	}
	return nil
}
