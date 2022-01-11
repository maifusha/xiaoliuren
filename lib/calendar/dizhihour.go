package calendar

var DizhiHours = map[int][]string{
	1:  {"子时", "23:00-01:00"},
	2:  {"丑时", "01:00-03:00"},
	3:  {"寅时", "03:00-05:00"},
	4:  {"卯时", "05:00-07:00"},
	5:  {"辰时", "07:00-09:00"},
	6:  {"巳时", "09:00-11:00"},
	7:  {"午时", "11:00-13:00"},
	8:  {"未时", "13:00-15;00"},
	9:  {"申时", "15:00-17:00"},
	10: {"酉时", "17:00-19:00"},
	11: {"戌时", "19:00-21:00"},
	12: {"亥时", "21:00-23:00"},
}

type dizhiHour struct {
	Index int
}

func NewDizhiHour(index int) *dizhiHour {
	return &dizhiHour{Index: index}
}

func (d *dizhiHour) Name() string {
	return DizhiHours[d.Index][0]
}

func (d *dizhiHour) Period() string {
	return DizhiHours[d.Index][1]
}
