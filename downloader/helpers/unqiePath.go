package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetUniqueFilePath(download string, filename string) string {
	var uniqueName string = filename
	counter := 1

	extension := filepath.Ext(uniqueName)
	name := strings.TrimSuffix(uniqueName, extension)

	for {
		_, err := os.Stat(filepath.Join(download,uniqueName))
		if err != nil{
			return uniqueName
		}
		uniqueName = fmt.Sprintf("%s(%d)%s",name,counter,extension)
		counter += 1
	}
}