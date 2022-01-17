package request

import (
	"time"

	"xiaoliuren/lib/liushen"
)

type Zeshi struct {
	Qike liushen.Gongwei `form:"qike" binding:"required,gte=1,lte=6"`
	Date time.Time       `form:"date" time_format:"2006-01-02" binding:"required"`
}

func NewZeshi() *Zeshi {
	return &Zeshi{}
}
