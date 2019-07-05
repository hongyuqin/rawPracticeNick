package models

type Score struct {
	Model
	CustomerId int `json:"customer_id"`
	Score      int `json:"score"`
}

func AddScore(score Score) error {
	if err := db.Create(&score).Error; err != nil {
		return err
	}
	return nil
}
