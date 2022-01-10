package model

import "gorm.io/gorm"

type qiuwen struct {
	gorm.Model
	LiushenId uint
	Type string
	Sentence string
}

func NewQiuwen() *qiuwen {
	return &qiuwen{}
}
