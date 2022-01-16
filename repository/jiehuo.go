package repository

import (
	"xiaoliuren/lib/liushen"
	"xiaoliuren/model"
)

type Jiehuo struct {
	model *model.Jiehuo
}

func NewJiehuo() *Jiehuo {
	return &Jiehuo{model: model.NewJiehuo()}
}

func (q *Jiehuo) SetModelById(id uint) {
	model := model.NewJiehuo()
	model.ID = id
	q.model = model
}

func (q *Jiehuo) FingByGongwei(gongwei liushen.Gongwei) (models []model.Jiehuo) {
	Db.Find(&models, "liushen_id", gongwei)
	return
}
