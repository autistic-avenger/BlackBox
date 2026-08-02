package filesio

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetPath() string {
	var filePath string
	osName := runtime.GOOS
	if osName == "windows"{
		appData := os.Getenv("LOCALAPPDATA")
		dirPath := filepath.Join(appData,"blackbox")
		filePath = filepath.Join(dirPath,"files.csv")
		_, err := os.Stat(filePath)
		if err!=nil{
			err := os.MkdirAll(dirPath,0755)
			if err!=nil{
				return ""
			}
			
			_, err = os.Create(filePath)
			if err!=nil{
				return ""
			}
		}
	}
	return filePath
}