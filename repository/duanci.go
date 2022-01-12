package repository

import (
	"xiaoliuren/lib/liushen"
	"xiaoliuren/model"
)

type Duanci struct {
	model *model.Duanci
}

func NewDuanci() *Duanci {
	return &Duanci{model: model.NewDuanci()}
}

func (d *Duanci) SetModelById(id uint) {
	model := model.NewDuanci()
	model.ID = id
	d.model = model
}

func (d *Duanci) FindByGongwei(gongwei liushen.Gongwei) (models []model.Duanci) {
	Db.Find(&models, "liushen_id=?", gongwei)
	return
}
