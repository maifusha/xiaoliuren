package request

import (
	"time"

	"xiaoliuren/lib/liushen"
)

type Zeshi struct {
	Qike liushen.Gongwei `json:"liushen_id" binding:"required,gte=1,lte=6"`
	Date time.Time       `json:"date" time_format:"2006-01-02" binding:"required"`
}

func NewZeshi() *Zeshi {
	return &Zeshi{}
}
