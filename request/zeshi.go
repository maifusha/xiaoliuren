package request

import "time"

type zeshi struct {
	LiushenId int       `json:"liushen_id" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
}

func NewZeshi() *zeshi {
	return &zeshi{}
}
