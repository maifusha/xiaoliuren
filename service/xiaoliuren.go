package service

import (
	"fmt"
	"sync"
	"time"

	"xiaoliuren/lib/calendar"
	"xiaoliuren/lib/entity"
	"xiaoliuren/repository"
)

type xiaoliuren struct {
}

func NewXiaoliuren() *xiaoliuren {
	return &xiaoliuren{}
}

func (l *xiaoliuren) GetLunarTime(date time.Time, dizhiIndex int) string {
	lunar := calendar.NewLunarBySolar(date)
	dizhiHour := calendar.NewDizhiHour(dizhiIndex)

	return fmt.Sprintf("%s%s", lunar.String(), dizhiHour.Name())
}

func (l *xiaoliuren) GetSolarTime(date time.Time, dizhiIndex int) string {
	dizhiHour := calendar.NewDizhiHour(dizhiIndex)

	return fmt.Sprintf("%s %s", date.Format("2006-01-02"), dizhiHour.Period())
}

func (l *xiaoliuren) GetSanPan(date time.Time, dizhiIndex int, liushenId int) (yuePan *entity.Jieke, riPan *entity.Jieke, shiPan *entity.Jieke) {
	lunar := calendar.NewLunarBySolar(date)
	dizhiHour := calendar.NewDizhiHour(dizhiIndex)

	month := lunar.MonthInt()
	day := lunar.DayInt()
	hour := dizhiHour.Index

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		yuePan = l.JieKe(liushenId, month)
	}()

	go func() {
		defer wg.Done()
		riPan = l.JieKe(liushenId, month+day-1)
	}()

	go func() {
		defer wg.Done()
		shiPan = l.JieKe(liushenId, month+day+hour-2)
	}()

	wg.Wait()
	return
}

func (l *xiaoliuren) JieKe(qike int, count int) *entity.Jieke {
	locate := (qike + count - 1) % 6
	liushen := repository.NewLiushen().FindById(locate)

	return entity.NewJieke(liushen.Name, liushen.Shiyi)
}
