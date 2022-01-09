package templatekit

import (
	"embed"
	"path"
	"strings"

	"github.com/gin-contrib/multitemplate"
)

type templateKit struct {
	fs *embed.FS
}

func New(fs *embed.FS) *templateKit {
	return &templateKit{fs: fs}
}

func (t *templateKit) MultiRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	var fragments string
	fragmentFiles, _ := t.fs.ReadDir("template/fragment")
	for _, fragmentFile := range fragmentFiles {
		fragmentFilePath := path.Join("template/fragment", fragmentFile.Name())
		fragment, _ := t.fs.ReadFile(fragmentFilePath)
		fragments += string(fragment)
	}

	contentDirs, _ := t.fs.ReadDir("template/content")
	for _, contentDir := range contentDirs {
		contentDirPath := path.Join("template/content", contentDir.Name())
		contentFiles, _ := t.fs.ReadDir(contentDirPath)

		for _, contentFile := range contentFiles {
			contentFilePath := path.Join(contentDirPath, contentFile.Name())
			tplName := strings.TrimPrefix(contentFilePath, "template/content/")
			content, _ := t.fs.ReadFile(contentFilePath)
			r.AddFromString(tplName, fragments+string(content))
		}
	}

	return r
}

