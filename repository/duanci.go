package repository

import "xiaoliuren/model"

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

func (d *Duanci) FindByLiushen(liushenId int) (models []model.Duanci) {
	Db.Find(&models, "liushen_id=?", liushenId)
	return
}
