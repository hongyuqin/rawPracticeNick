package models

type Score struct {
	Model
	CustomerId int `json:"customer_id"`
	Score      int `json:"score"`
}
