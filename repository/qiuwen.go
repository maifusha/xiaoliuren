package repository

import (
	"xiaoliuren/lib/liushen"
	"xiaoliuren/model"
)

type Qiuwen struct {
	model *model.Qiuwen
}

func newQiuwen() *Qiuwen {
	return &Qiuwen{model: model.NewQiuwen()}
}

func (q *Qiuwen) SetModelById(id uint) {
	model := model.NewQiuwen()
	model.ID = id
	q.model = model
}

func (q *Qiuwen) FingByGongwei(gongwei liushen.Gongwei) (models []model.Qiuwen) {
	Db.Find(&models, "liushen_id", gongwei)
	return
}
