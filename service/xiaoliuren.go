package service

import (
	"fmt"
	"sync"
	"time"

	"xiaoliuren/lib/calendar"
	"xiaoliuren/lib/liushen"
	"xiaoliuren/model"
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

	return fmt.Sprintf("%s %s", lunar.String(), dizhiHour.Name())
}

func (x *xiaoliuren) GetSolarTime(date time.Time, dizhi calendar.Dizhi) string {
	dizhiHour := calendar.NewDizhiHour(dizhi)

	return fmt.Sprintf("%s %s", date.Format("2006-01-02"), dizhiHour.Period())
}

func (x *xiaoliuren) LiushenList() []model.Liushen {
	qikeList := repository.NewLiushen().FindAll()

	return qikeList
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
	gongwei := liushen.LuogongByCount(qike, count)
	model := repository.NewLiushen().FindById(int(gongwei))

	return liushen.NewJieke(gongwei, model.Name, model.Shiyi)
}

func (x *xiaoliuren) GetShengong(gongwei liushen.Gongwei) *model.Liushen {
	shengong := repository.NewLiushen().FindById(int(gongwei))

	return shengong
}

func (x *xiaoliuren) JiehuoList(gongwei liushen.Gongwei) []model.Jiehuo {
	jiehuoList := repository.NewJiehuo().FingByGongwei(gongwei)

	return jiehuoList
}

func (x *xiaoliuren) DuanciList(gongwei liushen.Gongwei) []model.Duanci {
	duanciList := repository.NewDuanci().FindByGongwei(gongwei)

	return duanciList
}

func (x *xiaoliuren) JudgeHoursForDate(qike liushen.Gongwei, date time.Time) (
	daanList []interface{},
	suxiList []interface{},
	xiaojiList []interface{},
	liulianList []interface{},
	chikouList []interface{},
	kongwangList []interface{},
) {
	for k, v := range calendar.DizhiHours {
		item := struct {
			DizhiName string `json:"dizhi_name"`
			SolarTime string `json:"solar_time"`
		}{v[0], fmt.Sprintf("%s %s", date.Format("2006-01-02"), v[1])}

		switch liushen.LuogongByTime(qike, date, k) {
		case liushen.DAAN:
			daanList = append(daanList, item)
		case liushen.LIULIAN:
			liulianList = append(liulianList, item)
		case liushen.SUXI:
			suxiList = append(suxiList, item)
		case liushen.CHIKOU:
			chikouList = append(chikouList, item)
		case liushen.XIAOJI:
			xiaojiList = append(xiaojiList, item)
		case liushen.KONGWANG:
			kongwangList = append(kongwangList, item)
		}
	}

	return
}
