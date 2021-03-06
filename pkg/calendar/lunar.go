package calendar

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	thirdlunar "github.com/FrankWong1213/golang-lunar"
)

var (
	lunarMonths = map[string]int{"正月": 1, "二月": 2, "三月": 3, "四月": 4, "五月": 5, "六月": 6, "七月": 7, "八月": 8, "九月": 9, "十月": 10, "十一月": 11, "腊月": 12}
	lunarDays   = map[string]int{
		"初一": 1, "初二": 2, "初三": 3, "初四": 4, "初五": 5, "初六": 6, "初七": 7, "初八": 8, "初九": 9, "初十": 10,
		"十一": 11, "十二": 12, "十三": 13, "十四": 14, "十五": 15, "十六": 16, "十七": 17, "十八": 18, "十九": 19, "二十": 20,
		"廿一": 21, "廿二": 22, "廿三": 23, "廿四": 24, "廿五": 25, "廿六": 26, "廿七": 27, "廿八": 28, "廿九": 29, "三十": 30,
	}
)

type Lunar struct {
	year, month, day string
}

func NewLunarBySolar(solar time.Time) *Lunar {
	lunarStr := thirdlunar.Lunar(solar.Format("20060102"))
	result := regexp.MustCompile(`(?P<year>.*)年(?P<month>.*)月(?P<day>.*)`).FindAllStringSubmatch(lunarStr, -1)
	year, month, day := result[0][1], result[0][2], result[0][3]

	return &Lunar{
		year:  string([]rune(year)[0:2]) + "年",
		month: strings.NewReplacer("一", "正", "十二", "腊").Replace(month) + "月",
		day:   strings.NewReplacer("廿十", "二十", "卅十", "三十").Replace(day),
	}
}

func (l *Lunar) MonthInt() int {
	return lunarMonths[l.month]
}

func (l *Lunar) DayInt() int {
	return lunarDays[l.day]
}

func (l *Lunar) String() string {
	return fmt.Sprintf("%s %s%s", l.year, l.month, l.day)
}
