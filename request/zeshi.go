package request

import "time"

type Zeshi struct {
	LiushenId int       `json:"liushen_id" binding:"required,gte=1,lte=6"`
	Date      time.Time `json:"date" binding:"required"`
}

func NewZeshi() *Zeshi {
	return &Zeshi{}
}
