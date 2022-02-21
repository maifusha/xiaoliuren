package router

import (
	"embed"
	"io/fs"
	"net/http"

	"xiaoliuren/internal/router/handler"
	"xiaoliuren/internal/util/logger"
	"xiaoliuren/pkg/templatekit"

	"github.com/gin-gonic/gin"
)

type router struct {
	*gin.Engine
}

func New(engine *gin.Engine) *router {
	return &router{Engine: engine}
}

func (r *router) SetRenderWithEmbed(templateFS *embed.FS) *router {
	r.HTMLRender = templatekit.New(templateFS).MultiRender()

	return r
}

func (r *router) StaticBind(staticFS *embed.FS) *router {
	subStatic, err := fs.Sub(staticFS, "static")
	if err != nil {
		logger.Fatalln(err)
	}

	r.StaticFS("/static", http.FS(subStatic))

	return r
}

func (r *router) HandleBind() *router {
	homeHandle := handler.NewHome()
	r.GET("/", homeHandle.Redirect)

	homeGroup := r.Group("/home")
	{
		homeGroup.GET("/index", homeHandle.Index)

		homeGroup.GET("/jixiong", homeHandle.Jixiong)

		homeGroup.GET("/dianbo", homeHandle.Dianbo)

		homeGroup.GET("/zeshi", homeHandle.Zeshi)
	}

	return r
}
