package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
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
		//必须在router初始化前配置才有效
		gin.DefaultWriter, _ = os.Create(os.Getenv("GOPATH") + config.RUNTIME_LOG)
		gin.DefaultErrorWriter, _ = os.Create(os.Getenv("GOPATH") + config.ERROR_LOG)
	}

	//从嵌入式文件系统加载模板和静态资源
	router = gin.Default()
	router.HTMLRender = templatekit.New(&templateFS).MultiRender()
	subStatic, _ := fs.Sub(staticFS, "static")
	router.StaticFS("/static", http.FS(subStatic))
}

func main() {
	router.GET("/home/index", func(c *gin.Context) {
		svc := service.NewXiaoliuren()

		c.HTML(http.StatusOK, "home/index.tpl", gin.H{
			"qikeList":  svc.LiushenList(),
			"dizhiList": calendar.DizhiHours,
		})
	})

	router.GET("/home/jixiong", func(c *gin.Context) {
		req := request.NewJixiong()
		if c.ShouldBind(req) != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		svc := service.NewXiaoliuren()
		yuePan, riPan, shiPan := svc.GetSanPan(req.Qike, req.Date, req.Dizhi)

		c.JSON(http.StatusOK, gin.H{
			"lunar_time": svc.GetLunarTime(req.Date, req.Dizhi),
			"solar_time": svc.GetSolarTime(req.Date, req.Dizhi),
			"sanpan": gin.H{
				"yue_pan": yuePan,
				"ri_pan":  riPan,
				"shi_pan": shiPan,
			},
		})
	})

	router.GET("/home/dianbo", func(c *gin.Context) {
		date := time.Now()
		dizhi := calendar.NowDizhi()
		gongwei := liushen.LuogongByTime(liushen.DAAN, date, dizhi)

		svc := service.NewXiaoliuren()
		var shengong *model.Liushen
		var qiuwenList []model.Qiuwen
		var duanciList []model.Duanci

		var wg sync.WaitGroup
		wg.Add(3)

		go func() {
			defer wg.Done()
			shengong = svc.GetShengong(gongwei)
		}()

		go func() {
			defer wg.Done()
			qiuwenList = svc.QiuwenList(gongwei)
		}()

		go func() {
			defer wg.Done()
			duanciList = svc.DuanciList(gongwei)
		}()

		wg.Wait()

		c.JSON(http.StatusOK, gin.H{
			"lunar_time":  svc.GetLunarTime(date, dizhi),
			"solar_time":  svc.GetSolarTime(date, dizhi),
			"shengong":    shengong,
			"qiuwen_list": qiuwenList,
			"duanci_list": duanciList,
		})
	})

	router.GET("/home/zeshi", func(c *gin.Context) {
		req := request.NewZeshi()
		if c.ShouldBind(req) != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		svc := service.NewXiaoliuren()
		daanList, liulianList, suxiList, chikouList, xiaojiList, kongwangList := svc.JudgeHoursForDate(req.Qike, req.Date)

		c.JSON(http.StatusOK, gin.H{
			"daan":     daanList,
			"lulian":   liulianList,
			"suxi":     suxiList,
			"chikou":   chikouList,
			"xiaoji":   xiaojiList,
			"kongwang": kongwangList,
		})
	})

	log.Fatal(router.Run(config.HOST + ":" + config.PORT))
}
