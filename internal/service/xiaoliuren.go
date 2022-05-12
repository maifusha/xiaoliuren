package service

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"xiaoliuren/internal/filter"
	"xiaoliuren/internal/model"
	"xiaoliuren/internal/repository"
	"xiaoliuren/internal/util/logger"
	"xiaoliuren/pkg/calendar"
	"xiaoliuren/pkg/liuren"
)

type xiaoliuren struct {
}

func NewXiaoliuren() *xiaoliuren {
	return &xiaoliuren{}
}

func (x *xiaoliuren) LiushenList() []model.Liushen {
	qikeList := repository.NewLiushen().FindAll()

	return qikeList
}

func (x *xiaoliuren) GetSanGong(qike liuren.Gongwei, lunar *calendar.Lunar, dizhi *calendar.Dizhi) (yueke *liuren.Jieke, rike *liuren.Jieke, shike *liuren.Jieke) {
	month := lunar.MonthInt()
	day := lunar.DayInt()
	hour := dizhi.HourInt()

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

func (x *xiaoliuren) JieKe(qike liuren.Gongwei, count int) *liuren.Jieke {
	gongwei := liuren.FingerByCount(qike, count)
	liushen, err := repository.NewLiushen().FindById(int(gongwei))
	if err != nil {
		logger.Fatalln(err)
	}

	return liuren.NewJieke(gongwei, liushen.Name, liushen.Jixiong, liushen.Shiyi)
}

func (x *xiaoliuren) GetLiushen(gongwei liuren.Gongwei) model.Liushen {
	liushen, err := repository.NewLiushen().FindById(int(gongwei))
	if err != nil {
		logger.Fatalln(err)
	}

	return liushen
}

func (x *xiaoliuren) JiehuoList(f *filter.Jiehuo) []model.Jiehuo {
	jiehuoList := repository.NewJiehuo().Find(f)

	return jiehuoList
}

func (x *xiaoliuren) DuanciList(f *filter.Duanci) []model.Duanci {
	duanciList := repository.NewDuanci().Find(f)

	return duanciList
}

func (x *xiaoliuren) JudgeHoursForDate(qike liuren.Gongwei, date time.Time) (
	daanList []interface{},
	liulianList []interface{},
	suxiList []interface{},
	chikouList []interface{},
	xiaojiList []interface{},
	kongwangList []interface{},
) {
	var hours []calendar.Hour
	for k := range calendar.DizhiHours {
		hours = append(hours, k)
	}
	sort.Slice(hours, func(i, j int) bool {
		return hours[i] < hours[j]
	})

	lunar := calendar.NewLunarBySolar(date)
	for _, hour := range hours {
		dizhi := calendar.NewDizhi(hour)
		item := struct {
			DizhiName string `json:"dizhi_name"`
			SolarTime string `json:"solar_time"`
		}{dizhi.Name(), fmt.Sprintf("%s %s", date.Format("2006-01-02"), dizhi.Period())}

		switch liuren.FingerByTime(qike, lunar, dizhi) {
		case liuren.DAAN:
			daanList = append(daanList, item)
		case liuren.LIULIAN:
			liulianList = append(liulianList, item)
		case liuren.SUXI:
			suxiList = append(suxiList, item)
		case liuren.CHIKOU:
			chikouList = append(chikouList, item)
		case liuren.XIAOJI:
			xiaojiList = append(xiaojiList, item)
		case liuren.KONGWANG:
			kongwangList = append(kongwangList, item)
		}
	}

	return
}
