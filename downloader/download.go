package downloader

import (
	"blackbox/downloader/networks"
	"blackbox/models"
)



func DownloadFile(file *models.File)  {
	err := networks.MakeRequest(file)
	if err!=nil{
		file.Mu.Lock()
		file.Error = "ERROR: http req failed :("
		file.Mu.Unlock()
	}
}