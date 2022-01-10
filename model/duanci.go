package model

import "gorm.io/gorm"

type duanci struct {
	gorm.Model
	LiushenId uint
	Sentence string
}

func NewDuanci() *duanci {
	return &duanci{}
}
