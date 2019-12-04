package files

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func Load(filePath string) []string {
	absPath, _ := filepath.Abs(filePath)
	dat, err := ioutil.ReadFile(absPath)

	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(dat), "\n")
}
