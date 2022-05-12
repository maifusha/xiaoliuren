package liuren

import (
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

func FingerByTime(qike Gongwei, lunar *calendar.Lunar, dizhi *calendar.Dizhi) Gongwei {
	month := lunar.MonthInt()
	day := lunar.DayInt()
	hour := dizhi.HourInt()

	return FingerByCount(qike, month+day+hour-2)
}

func FingerByCount(qike Gongwei, count int) Gongwei {
	luogon := (int(qike) + count - 1) % 6
	if luogon == 0 {
		luogon = 6
	}

	return Gongwei(luogon)
}
