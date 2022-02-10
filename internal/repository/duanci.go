package repository

import (
	"gorm.io/gorm"

	"xiaoliuren/internal/filter"
	"xiaoliuren/internal/model"
)

type Duanci struct {
	model *model.Duanci
}

func NewDuanci() *Duanci {
	return &Duanci{model: model.NewDuanci()}
}

func (d *Duanci) Find(f *filter.Duanci) (models []model.Duanci) {
	query := Db.Session(&gorm.Session{})

	if f.LiushenId != 0 {
		query.Where("liushen_id=?", f.LiushenId)
	}

	query.Find(&models)

	return
}
