package model

import (
	"gorm.io/gorm"
)

type Liushen struct {
	gorm.Model
	Name           string `json:"name"`
	Jixiong        string `json:"jixiong"`
	Shensha        string `json:"shensha"`
	Wuxin          string `json:"wuxin"`
	Bagua          string `json:"bagua"`
	Shiergong      string `json:"shiergong"`
	Paiweishu      string `json:"paiweishu"`
	Gongweishu     string `json:"gongweishu"`
	Fangwei        string `json:"fangwei"`
	Guirenchongfan string `json:"guirenchongfan"`
	Suozhu         string `json:"suozhu"`
	Shiyi          string `json:"shiyi"`
}

func NewLiushen() *Liushen {
	return &Liushen{}
}
