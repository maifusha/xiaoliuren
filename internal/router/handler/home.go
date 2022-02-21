package handler

import (
	"net/http"
	"sync"
	"time"

	"xiaoliuren/internal/filter"
	"xiaoliuren/internal/model"
	"xiaoliuren/internal/request"
	"xiaoliuren/internal/service"
	"xiaoliuren/pkg/calendar"
	"xiaoliuren/pkg/liushen"

	"github.com/gin-gonic/gin"
)

type Home struct {
}

func NewHome() *Home {
	return &Home{}
}

func (h *Home) Redirect(c *gin.Context) {
	c.Redirect(http.StatusFound, "/home/index")
}

func (h *Home) Index(c *gin.Context) {
	xlr := service.NewXiaoliuren()

	c.HTML(http.StatusOK, "home/index.tpl", gin.H{
		"qikeList":  xlr.LiushenList(),
		"dizhiList": calendar.DizhiHours,
	})
}

func (h *Home) Jixiong(c *gin.Context) {
	req := request.NewJixiong()
	if c.ShouldBind(req) != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	xlr := service.NewXiaoliuren()
	yueke, rike, shike := xlr.GetSanGong(req.Qike, req.Date, req.Dizhi)

	lunar := calendar.NewLunarBySolar(req.Date)

	c.JSON(http.StatusOK, gin.H{
		"qike":         req.Qike,
		"finger_count": lunar.MonthInt() + lunar.DayInt() + int(req.Dizhi) - 2,
		"lunar_time":   xlr.GetLunarTime(req.Date, req.Dizhi),
		"solar_time":   xlr.GetSolarTime(req.Date, req.Dizhi),
		"sangong": gin.H{
			"yueke": yueke,
			"rike":  rike,
			"shike": shike,
		},
	})
}

func (h *Home) Dianbo(c *gin.Context) {
	date := time.Now()
	lunar := calendar.NewLunarBySolar(date)
	dizhi := calendar.NowDizhi()
	gongwei := liushen.LuogongByTime(liushen.DAAN, date, dizhi)

	xlr := service.NewXiaoliuren()
	var shengong model.Liushen
	var jiehuoList []model.Jiehuo
	var duanciList []model.Duanci

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		shengong = xlr.GetShengong(gongwei)
	}()

	go func() {
		defer wg.Done()
		f := filter.NewJiehuo()
		f.LiushenId = uint(gongwei)
		jiehuoList = xlr.JiehuoList(f)
	}()

	go func() {
		defer wg.Done()
		f := filter.NewDuanci()
		f.LiushenId = uint(gongwei)
		duanciList = xlr.DuanciList(f)
	}()

	wg.Wait()

	c.JSON(http.StatusOK, gin.H{
		"qike":         liushen.DAAN,
		"finger_count": lunar.MonthInt() + lunar.DayInt() + int(dizhi) - 2,
		"lunar_time":   xlr.GetLunarTime(date, dizhi),
		"solar_time":   xlr.GetSolarTime(date, dizhi),
		"shengong":     shengong,
		"jiehuo_list":  jiehuoList,
		"duanci_list":  duanciList,
	})
}

func (h *Home) Zeshi(c *gin.Context) {
	req := request.NewZeshi()
	if c.ShouldBind(req) != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	xlr := service.NewXiaoliuren()
	daanList, liulianList, suxiList, chikouList, xiaojiList, kongwangList := xlr.JudgeHoursForDate(req.Qike, req.Date)

	c.JSON(http.StatusOK, gin.H{
		"solar": req.Date.Format("2006-01-02"),
		"lunar": calendar.NewLunarBySolar(req.Date).String(),
		"liushen": gin.H{
			"daan":     daanList,
			"liulian":  liulianList,
			"suxi":     suxiList,
			"chikou":   chikouList,
			"xiaoji":   xiaojiList,
			"kongwang": kongwangList,
		},
	})
}
