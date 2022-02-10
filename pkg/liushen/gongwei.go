package liushen

import (
	"time"

	"xiaoliuren/pkg/calendar"
)

type Gongwei int

const (
	_ Gongwei = iota
	DAAN
	LIULIAN
	SUXI
	CHIKOU
	XIAOJI
	KONGWANG
)

func LuogongByTime(qike Gongwei, date time.Time, dizhi calendar.Dizhi) Gongwei {
	lunar := calendar.NewLunarBySolar(date)
	dizhiHour := calendar.NewDizhiHour(dizhi)

	month := lunar.MonthInt()
	day := lunar.DayInt()
	hour := int(dizhiHour.Dizhi)

	return LuogongByCount(qike, month+day+hour-2)
}

func LuogongByCount(qike Gongwei, count int) Gongwei {
	luogon := (int(qike) + count - 1) % 6
	if luogon == 0 {
		luogon = 6
	}

	return Gongwei(luogon)
}
