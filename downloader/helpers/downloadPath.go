package helpers

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetDownloadPath() string {
	osName := runtime.GOOS
	if osName == "windows"{
		homeDir,err := os.UserHomeDir()
		if err!=nil{
			return ""
		}
		return filepath.Join(homeDir,"Downloads")
	}
	return ""
}