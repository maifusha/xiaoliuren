package service

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"xiaoliuren/internal/filter"
	"xiaoliuren/internal/model"
	"xiaoliuren/internal/repository"
	"xiaoliuren/pkg/calendar"
	"xiaoliuren/pkg/liushen"
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
	qikeList, _ := repository.NewLiushen().FindAll()

	return qikeList
}

func (x *xiaoliuren) GetSanGong(qike liushen.Gongwei, date time.Time, dizhi calendar.Dizhi) (yueke *liushen.Jieke, rike *liushen.Jieke, shike *liushen.Jieke) {
	lunar := calendar.NewLunarBySolar(date)
	dizhiHour := calendar.NewDizhiHour(dizhi)

	month := lunar.MonthInt()
	day := lunar.DayInt()
	hour := int(dizhiHour.Dizhi)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		yueke = x.JieKe(qike, month)
	}()

	go func() {
		defer wg.Done()
		rike = x.JieKe(qike, month+day-1)
	}()

	go func() {
		defer wg.Done()
		shike = x.JieKe(qike, month+day+hour-2)
	}()

	wg.Wait()
	return
}

func (x *xiaoliuren) JieKe(qike liushen.Gongwei, count int) *liushen.Jieke {
	gongwei := liushen.LuogongByCount(qike, count)
	model, _ := repository.NewLiushen().FindById(int(gongwei))

	return liushen.NewJieke(gongwei, model.Name, model.Jixiong, model.Shiyi)
}

func (x *xiaoliuren) GetShengong(gongwei liushen.Gongwei) *model.Liushen {
	shengong, _ := repository.NewLiushen().FindById(int(gongwei))

	return shengong
}

func (x *xiaoliuren) JiehuoList(f *filter.Jiehuo) []model.Jiehuo {
	jiehuoList := repository.NewJiehuo().Find(f)

	return jiehuoList
}

func (x *xiaoliuren) DuanciList(f *filter.Duanci) []model.Duanci {
	duanciList := repository.NewDuanci().Find(f)

	return duanciList
}

func (x *xiaoliuren) JudgeHoursForDate(qike liushen.Gongwei, date time.Time) (
	daanList []interface{},
	liulianList []interface{},
	suxiList []interface{},
	chikouList []interface{},
	xiaojiList []interface{},
	kongwangList []interface{},
) {
	var dizhis []calendar.Dizhi
	for k := range calendar.DizhiHours {
		dizhis = append(dizhis, k)
	}
	sort.Slice(dizhis, func(i, j int) bool {
		return dizhis[i] < dizhis[j]
	})

	for _, dizhi := range dizhis {
		dizhiHour := calendar.DizhiHours[dizhi]
		item := struct {
			DizhiName string `json:"dizhi_name"`
			SolarTime string `json:"solar_time"`
		}{dizhiHour[0], fmt.Sprintf("%s %s", date.Format("2006-01-02"), dizhiHour[1])}

		switch liushen.LuogongByTime(qike, date, dizhi) {
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