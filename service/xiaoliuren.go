package service

import (
	"fmt"
	"sync"
	"time"

	"xiaoliuren/lib/calendar"
	"xiaoliuren/lib/liushen"
	"xiaoliuren/repository"
)

type xiaoliuren struct {
}

func NewXiaoliuren() *xiaoliuren {
	return &xiaoliuren{}
}

func (x *xiaoliuren) GetLunarTime(date time.Time, dizhi calendar.Dizhi) string {
	lunar := calendar.NewLunarBySolar(date)
	dizhiHour := calendar.NewDizhiHour(dizhi)

	return fmt.Sprintf("%s%s", lunar.String(), dizhiHour.Name())
}

func (x *xiaoliuren) GetSolarTime(date time.Time, dizhi calendar.Dizhi) string {
	dizhiHour := calendar.NewDizhiHour(dizhi)

	return fmt.Sprintf("%s %s", date.Format("2006-01-02"), dizhiHour.Period())
}

func (x *xiaoliuren) GetSanPan(qike liushen.Gongwei, date time.Time, dizhi calendar.Dizhi) (yuePan *liushen.Jieke, riPan *liushen.Jieke, shiPan *liushen.Jieke) {
	lunar := calendar.NewLunarBySolar(date)
	dizhiHour := calendar.NewDizhiHour(dizhi)

	month := lunar.MonthInt()
	day := lunar.DayInt()
	hour := int(dizhiHour.Dizhi)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		yuePan = x.JieKe(qike, month)
	}()

	go func() {
		defer wg.Done()
		riPan = x.JieKe(qike, month+day-1)
	}()

	go func() {
		defer wg.Done()
		shiPan = x.JieKe(qike, month+day+hour-2)
	}()

	wg.Wait()
	return
}

func (x *xiaoliuren) JieKe(qike liushen.Gongwei, count int) *liushen.Jieke {
	luogong := liushen.LuogongByCount(qike, count)
	model := repository.NewLiushen().FindById(int(luogong))

	return liushen.NewJieke(luogong, model.Name, model.Shiyi)
}
