package request

import (
	"time"

	"xiaoliuren/pkg/calendar"
	"xiaoliuren/pkg/liuren"
)

type Jixiong struct {
	Qike  liuren.Gongwei `form:"qike" binding:"required,gte=1,lte=6"`
	Date  time.Time      `form:"date" time_format:"2006-01-02" binding:"required"`
	Dizhi calendar.Hour  `form:"dizhi"  binding:"required,gte=1,lte=12"`
}

func NewJixiong() *Jixiong {
	return &Jixiong{}
}
