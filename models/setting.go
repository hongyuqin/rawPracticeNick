package models

type Setting struct {
	Model
	OpenId         string `json:"open_id"`
	ElementTypeOne string `json:"element_type_one"`
	ElementTypeTwo string `json:"element_type_two"`
}
