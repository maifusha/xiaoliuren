package model

import (
	"gorm.io/gorm"
)

type liushen struct {
	gorm.Model
	Name string
	Jixiong string
	Shensha string
	Wuxin string
	Bagua string
	Shiergong string
	Paiweishu string
	Gongweishu string
	Fangwei string
	Guirenchongfan string
	Suozhu string
	Shiyi string
}

func NewLiushen() *liushen {
	return &liushen{}
}
