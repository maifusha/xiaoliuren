package model

import "gorm.io/gorm"

type Qiuwen struct {
	gorm.Model
	LiushenId uint   `json:"liushen_id"`
	Type      string `json:"type"`
	Sentence  string `json:"sentence"`
}

func NewQiuwen() *Qiuwen {
	return &Qiuwen{}
}
