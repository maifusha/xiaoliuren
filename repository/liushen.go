package repository

import "xiaoliuren/model"

type Liushen struct {
	model *model.Liushen
}

func NewLiushen() *Liushen {
	return &Liushen{model: model.NewLiushen()}
}

func (l *Liushen) SetModelById(id uint) {
	model := model.NewLiushen()
	model.ID = id
	l.model = model
}

func (l *Liushen) FindById(id int) *model.Liushen {
	Db.First(l.model, id)

	return l.model
}

func (l *Liushen) FindAll() (models []model.Liushen) {
	Db.Find(&models)

	return
}
