package request

import (
	"time"

	"xiaoliuren/lib/calendar"
	"xiaoliuren/lib/liushen"
)

type Jixiong struct {
	Qike  liushen.Gongwei `form:"qike" binding:"required,gte=1,lte=6"`
	Date  time.Time       `form:"date" time_format:"2006-01-02" binding:"required"`
	Dizhi calendar.Dizhi  `form:"dizhi"  binding:"required,gte=1,lte=12"`
}

func NewJixiong() *Jixiong {
	return &Jixiong{}
}
