package request

import (
	"time"

	"xiaoliuren/lib/calendar"
	"xiaoliuren/lib/liushen"
)

type Jixiong struct {
	Qike  liushen.Gongwei `json:"liushen_id" binding:"required,gte=1,lte=6"`
	Date  time.Time       `json:"date" time_format:"2006-01-02" binding:"required"`
	Dizhi calendar.Dizhi  `json:"dizhi_hour"  binding:"required,gte=1,lte=12"`
}

func NewJixiong() *Jixiong {
	return &Jixiong{}
}
