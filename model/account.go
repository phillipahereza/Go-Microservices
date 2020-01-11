package model

type Account struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ServedBy string `json:"served_by"`
	Quote    Quote  `json:"quote"`
}

type Quote struct {
	Text     string `json:"quote"`
	ServedBy string `json:"ipAddress"`
	Language string `json:"language"`
}
