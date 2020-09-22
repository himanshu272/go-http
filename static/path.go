package static

import (
	"os"
	"log"
	"path/filepath"
)

func GetPath() string {
	dir, err := os.Getwd()

	if err != nil {
		log.Printf("Error!")
	}

	dirpath, _ := filepath.Abs(dir)

	return dirpath
}
