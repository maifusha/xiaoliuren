package model

import "gorm.io/gorm"

type Qiuwen struct {
	gorm.Model
	LiushenId uint
	Type      string
	Sentence  string
}

func NewQiuwen() *Qiuwen {
	return &Qiuwen{}
}
