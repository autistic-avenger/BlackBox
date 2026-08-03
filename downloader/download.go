package downloader

import (
	"blackbox/downloader/networks"
	"blackbox/models"
)



func DownloadFile(file *models.File)  {
	go networks.MakeRequest(file)
}