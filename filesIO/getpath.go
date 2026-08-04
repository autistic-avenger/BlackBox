package filesio

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetPath() string {
	var filePath string
	var dirPath string


	switch runtime.GOOS {
	case "windows":
		appData := os.Getenv("LOCALAPPDATA")
		dirPath = filepath.Join(appData, "blackbox")
	case "linux":
		configDir, err := os.UserConfigDir()
		if err != nil {
			return ""
		}
		dirPath = filepath.Join(configDir, "blackbox")
	case "darwin":
		configDir, err := os.UserConfigDir()
		if err != nil {
			return ""
		}
		dirPath = filepath.Join(configDir, "blackbox")

	default:
		return ""
	}

	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return ""
	}

	filePath = filepath.Join(dirPath, "files.csv")
	_, err = os.Stat(filePath)
	if err != nil {
		_, err = os.Create(filePath)
		if err != nil {
			return ""
		}
	}
	return filePath
}