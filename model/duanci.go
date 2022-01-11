package model

import "gorm.io/gorm"

type Duanci struct {
	gorm.Model
	LiushenId uint
	Sentence  string
}

func NewDuanci() *Duanci {
	return &Duanci{}
}
