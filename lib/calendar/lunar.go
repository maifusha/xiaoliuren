package calendar

import (
	"fmt"
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

type lunar struct {
	Year, Month, Day string
}

func NewLunarBySolar(solar time.Time) *lunar {
	year := ""
	month := ""
	day := ""
	lunarStr := thirdlunar.Lunar(solar.Format("20060102"))
	fmt.Scanf(lunarStr, "%s年%s月%s", &year, &month, &day)

	return &lunar{
		Year:  string([]rune(year)[0:2]) + "年",
		Month: strings.NewReplacer("一", "正", "十二", "腊").Replace(month) + "月",
		Day:   strings.NewReplacer("廿十", "二十", "卅十", "三十").Replace(day),
	}
}

func (l *lunar) MonthInt() int {
	return lunarMonths[l.Month]
}

func (l *lunar) DayInt() int {
	return lunarDays[l.Day]
}

func (l *lunar) String() string {
	return l.Year + l.Month + l.Day
}
