package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-contrib/multitemplate"
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
		gin.DefaultWriter, _ = os.Create(os.Getenv("GOPATH") + "/gin_runtime.log")
		gin.DefaultErrorWriter, _ = os.Create(os.Getenv("GOPATH") + "/gin_error.log")
	}

	router = gin.Default()

	//从嵌入式文件系统加载模板和静态资源
	router.HTMLRender = loadTemplates()
	subStatic, _ := fs.Sub(staticFS, "static")
	router.StaticFS("/static", http.FS(subStatic))
}

func main() {
	router.GET("/home/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.tpl", gin.H{})
	})

	log.Fatal(router.Run(":8000"))
}

func loadTemplates() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	var fragments string
	fragmentFiles, _ := templateFS.ReadDir("template/fragment")
	for _, fragmentFile := range fragmentFiles {
		fragmentFilePath := path.Join("template/fragment", fragmentFile.Name())
		fragment, _ := templateFS.ReadFile(fragmentFilePath)
		fragments += string(fragment)
	}

	contentDirs, _ := templateFS.ReadDir("template/content")
	for _, contentDir := range contentDirs {
		contentDirPath := path.Join("template/content", contentDir.Name())
		contentFiles, _ := templateFS.ReadDir(contentDirPath)

		for _, contentFile := range contentFiles {
			contentFilePath := path.Join(contentDirPath, contentFile.Name())
			tplName := strings.TrimPrefix(contentFilePath, "template/content/")
			content, _ := templateFS.ReadFile(contentFilePath)
			r.AddFromString(tplName, fragments+string(content))
		}
	}

	return r
}
