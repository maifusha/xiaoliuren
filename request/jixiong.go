package request

import "time"

type jixiong struct {
	LiushenId int       `json:"liushen_id" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	DizhiHour int       `json:"dizhi_hour" time_format:"2006-01-02" binding:"required"`
}

func NewJixiong() *jixiong {
	return &jixiong{}
}
