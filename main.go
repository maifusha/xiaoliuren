package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
	"time"

	"xiaoliuren/config"
	"xiaoliuren/lib/calendar"
	"xiaoliuren/lib/liushen"
	"xiaoliuren/lib/templatekit"
	"xiaoliuren/model"
	"xiaoliuren/request"
	"xiaoliuren/service"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

//go:embed template
var templateFS embed.FS

//go:embed static
var staticFS embed.FS

func init() {
	if gin.Mode() == gin.ReleaseMode {
		gin.DefaultWriter, _ = os.Create(path.Join(os.Getenv("GOPATH"), config.RUNTIME_LOG))
		gin.DefaultErrorWriter, _ = os.Create(path.Join(os.Getenv("GOPATH"), config.ERROR_LOG))
	}

	router = gin.Default()
	router.HTMLRender = templatekit.New(&templateFS).MultiRender()
	subStatic, _ := fs.Sub(staticFS, "static")
	router.StaticFS("/static", http.FS(subStatic))
}

func main() {
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/home/index")
	})

	router.GET("/home/index", func(c *gin.Context) {
		xlr := service.NewXiaoliuren()

		c.HTML(http.StatusOK, "home/index.tpl", gin.H{
			"qikeList":  xlr.LiushenList(),
			"dizhiList": calendar.DizhiHours,
		})
	})

	router.GET("/home/jixiong", func(c *gin.Context) {
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
	})

	router.GET("/home/dianbo", func(c *gin.Context) {
		date := time.Now()
		lunar := calendar.NewLunarBySolar(date)
		dizhi := calendar.NowDizhi()
		gongwei := liushen.LuogongByTime(liushen.DAAN, date, dizhi)

		xlr := service.NewXiaoliuren()
		var shengong *model.Liushen
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
			jiehuoList = xlr.JiehuoList(gongwei)
		}()

		go func() {
			defer wg.Done()
			duanciList = xlr.DuanciList(gongwei)
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
	})

	router.GET("/home/zeshi", func(c *gin.Context) {
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
	})

	log.Fatal(router.Run(config.HOST + ":" + config.PORT))
}
