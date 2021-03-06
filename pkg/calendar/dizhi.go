package calendar

import (
	"strings"
	"time"
)

type Hour int

const (
	_ Hour = iota
	ZISHI
	CHOUSHI
	YINSHI
	MOUSHI
	CHENSHI
	SISHI
	WUSHI
	WEISHI
	SHENSHI
	YOUSHI
	XUSHI
	HAISHI
)

var DizhiHours = map[Hour][]string{
	ZISHI:   {"子时", "23:00~01:00"},
	CHOUSHI: {"丑时", "01:00~03:00"},
	YINSHI:  {"寅时", "03:00~05:00"},
	MOUSHI:  {"卯时", "05:00~07:00"},
	CHENSHI: {"辰时", "07:00~09:00"},
	SISHI:   {"巳时", "09:00~11:00"},
	WUSHI:   {"午时", "11:00~13:00"},
	WEISHI:  {"未时", "13:00~15;00"},
	SHENSHI: {"申时", "15:00~17:00"},
	YOUSHI:  {"酉时", "17:00~19:00"},
	XUSHI:   {"戌时", "19:00~21:00"},
	HAISHI:  {"亥时", "21:00~23:00"},
}

type Dizhi struct {
	Hour
}

func NewDizhi(hour Hour) *Dizhi {
	return &Dizhi{Hour: hour}
}

func (d *Dizhi) HourInt() int {
	return int(d.Hour)
}

func (d *Dizhi) Name() string {
	return DizhiHours[d.Hour][0]
}

func (d *Dizhi) Period() string {
	return DizhiHours[d.Hour][1]
}

func NowDizhi() *Dizhi {
	now := time.Now().Format("15:04")

	var hour Hour
	for k, v := range DizhiHours {
		period := strings.Split(v[1], "~")
		if strings.Compare(now, period[0]) >= 0 && strings.Compare(now, period[1]) < 0 {
			hour = k
		}
	}

	return NewDizhi(hour)
}
