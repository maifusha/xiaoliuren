package file

import (
	"log"
	"os"
	"path"
)

func MustOpen(filepath string) *os.File {
	if err := os.MkdirAll(path.Dir(filepath), 0777); err != nil {
		log.Fatalln(err)
	}

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	return f
}
