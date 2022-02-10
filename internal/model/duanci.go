package model

import "gorm.io/gorm"

type Duanci struct {
	gorm.Model
	LiushenId uint   `json:"liushen_id"`
	Sentence  string `json:"sentence"`
}

func NewDuanci() *Duanci {
	return &Duanci{}
}
