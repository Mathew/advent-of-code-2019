package files

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func Load(filePath string, separator string) []string {
	absPath, _ := filepath.Abs(filePath)
	dat, err := ioutil.ReadFile(absPath)

	if err != nil {
		log.Fatal(err)
	}

	if separator == "" {
		return []string{strings.TrimSpace(string(dat))}
	}

	return strings.Split(string(dat), separator)
}
