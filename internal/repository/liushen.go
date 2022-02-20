package repository

import "xiaoliuren/internal/model"

type Liushen struct {
}

func NewLiushen() *Liushen {
	return &Liushen{}
}

func (l *Liushen) FindById(id int) (m model.Liushen, err error) {
	err = Db.First(&m, id).Error

	return
}

func (l *Liushen) FindAll() (models []model.Liushen) {
	Db.Find(&models)

	return
}
