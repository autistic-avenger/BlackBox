package downloader

import (
	"github.com/autistic-avenger/BlackBox/downloader/networks"
	"github.com/autistic-avenger/BlackBox/models"
)



func DownloadFile(file *models.File)  {
	err := networks.MakeRequest(file)
	if err!=nil{
		file.Mu.Lock()
		file.Error = "ERROR: http req failed :("
		file.Mu.Unlock()
	}
}