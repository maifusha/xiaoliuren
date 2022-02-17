package repository

import (
	"gorm.io/gorm"

	"xiaoliuren/internal/filter"
	"xiaoliuren/internal/model"
)

type Jiehuo struct {
	model *model.Jiehuo
}

func NewJiehuo() *Jiehuo {
	return &Jiehuo{model: model.NewJiehuo()}
}

func (q *Jiehuo) Find(f *filter.Jiehuo) (models []model.Jiehuo) {
	query := Db.Session(&gorm.Session{})

	if f.LiushenId != 0 {
		query = query.Where("liushen_id=?", f.LiushenId)
	}

	query.Find(&models)

	return
}
