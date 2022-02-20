package repository

import (
	"gorm.io/gorm"

	"xiaoliuren/internal/filter"
	"xiaoliuren/internal/model"
)

type Duanci struct {
}

func NewDuanci() *Duanci {
	return &Duanci{}
}

func (d *Duanci) Find(f *filter.Duanci) (models []model.Duanci) {
	query := Db.Session(&gorm.Session{})

	if f.LiushenId != 0 {
		query = query.Where("liushen_id=?", f.LiushenId)
	}

	query.Find(&models)

	return
}
