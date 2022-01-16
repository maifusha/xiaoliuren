package model

import "gorm.io/gorm"

type Jiehuo struct {
	gorm.Model
	LiushenId uint   `json:"liushen_id"`
	Type      string `json:"type"`
	Sentence  string `json:"sentence"`
}

func NewJiehuo() *Jiehuo {
	return &Jiehuo{}
}
