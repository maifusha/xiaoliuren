package request

import "time"

type Jixiong struct {
	LiushenId  int       `json:"liushen_id" binding:"required,gte=1,lte=6"`
	Date       time.Time `json:"date" binding:"required"`
	DizhiIndex int       `json:"dizhi_hour" time_format:"2006-01-02" binding:"required,gte=1,lte=12"`
}

func NewJixiong() *Jixiong {
	return &Jixiong{}
}
