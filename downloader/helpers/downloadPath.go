package helpers

import (
	"os"
	"path/filepath"
)

func GetDownloadPath() string {
	homeDir,err := os.UserHomeDir()
	if err!=nil{
		return ""
	}
	return filepath.Join(homeDir,"Downloads")
}