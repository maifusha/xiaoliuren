package home

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"xiaoliuren/internal/filter"
	"xiaoliuren/internal/model"
	"xiaoliuren/internal/request"
	"xiaoliuren/internal/service"
	"xiaoliuren/pkg/calendar"
	"xiaoliuren/pkg/liuren"

	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	c.Redirect(http.StatusFound, "/home/index")
}

func Index(c *gin.Context) {
	xlr := service.NewXiaoliuren()

	c.HTML(http.StatusOK, "home/index.tpl", gin.H{
		"qikeList":  xlr.LiushenList(),
		"dizhiList": calendar.DizhiHours,
	})
}

func Jixiong(c *gin.Context) {
	req := request.NewJixiong()
	if c.ShouldBind(req) != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	lunar := calendar.NewLunarBySolar(req.Date)
	dizhi := calendar.NewDizhi(req.Dizhi)

	xlr := service.NewXiaoliuren()
	yueke, rike, shike := xlr.GetSanGong(req.Qike, lunar, dizhi)

	c.JSON(http.StatusOK, gin.H{
		"qike":         req.Qike,
		"finger_count": lunar.MonthInt() + lunar.DayInt() + dizhi.HourInt() - 2,
		"lunar_time":   fmt.Sprintf("%s %s", lunar.String(), dizhi.Name()),
		"solar_time":   fmt.Sprintf("%s %s", req.Date.Format("2006-01-02"), dizhi.Period()),
		"sangong": gin.H{
			"yueke": yueke,
			"rike":  rike,
			"shike": shike,
		},
	})
}

func Dianbo(c *gin.Context) {
	date := time.Now()
	lunar := calendar.NewLunarBySolar(date)
	dizhi := calendar.NowDizhi()
	gongwei := liuren.FingerByTime(liuren.DAAN, lunar, dizhi)

	xlr := service.NewXiaoliuren()
	var liushen model.Liushen
	var jiehuoList []model.Jiehuo
	var duanciList []model.Duanci

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		liushen = xlr.GetLiushen(gongwei)
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
		"qike":         liuren.DAAN,
		"finger_count": lunar.MonthInt() + lunar.DayInt() + dizhi.HourInt() - 2,
		"lunar_time":   fmt.Sprintf("%s %s", lunar.String(), dizhi.Name()),
		"solar_time":   fmt.Sprintf("%s %s", date.Format("2006-01-02"), dizhi.Period()),
		"liushen":      liushen,
		"jiehuo_list":  jiehuoList,
		"duanci_list":  duanciList,
	})
}

func Zeshi(c *gin.Context) {
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
