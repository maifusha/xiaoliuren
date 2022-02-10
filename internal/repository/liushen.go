package repository

import "xiaoliuren/internal/model"

type Liushen struct {
	model *model.Liushen
}

func NewLiushen() *Liushen {
	return &Liushen{model: model.NewLiushen()}
}

func (l *Liushen) FindById(id int) (*model.Liushen, error) {
	result := Db.First(l.model, id)

	return l.model, result.Error
}

func (l *Liushen) FindAll() (models []model.Liushen, cnt int64) {
	result := Db.Find(&models)
	cnt = result.RowsAffected

	return
}
