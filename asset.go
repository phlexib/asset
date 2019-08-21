package asset

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

//GetFileVersion uses Regex to finf the file version
func GetFileVersion(str string) int {
	var version = regexp.MustCompile(`^(.*?)[vV]+([0-9]{1,})`)
	// baseName := filepath.Base(str)
	res := version.FindStringSubmatch(str)
	v, err := strconv.Atoi(res[2])
	if err != nil {
		log.Fatal("No version found")
	}
	return v
}

//UpFileName find a version Number in a filename and return the next one adjust padding as following "%03d"
func UpFileName(str string, padding string) string {
	newVersion := GetFileVersion(str) + 1
	vString := "v" + padNumberWithZero(padding, newVersion)
	var vReg = regexp.MustCompile(`[vV]+[0-9]{1,}`)
	return vReg.ReplaceAllString(str, vString)
}

func padNumberWithZero(padding string, value int) string {
	return fmt.Sprintf(padding, value)
}

func getFilesFromPath(str string, search string, ignore []string, showArchive bool) []string {
	dir := filepath.Dir(str)
	if !showArchive {
		ignore = append(ignore, "_archive")
	}
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	r := regexp.MustCompile(search)

	filtered := Filter(files, func(v string) bool {
		return r.MatchString(v)
	})

	if len(ignore) > 0 {
		for i := 0; i < len(ignore); i++ {
			filtered = Filter(filtered, func(v string) bool {
				return !strings.Contains(v, ignore[i])
			})
		}
	}

	return filtered
}
