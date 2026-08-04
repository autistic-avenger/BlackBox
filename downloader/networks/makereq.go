package networks

import (
	"blackbox/downloader/helpers"
	"blackbox/models"
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func MakeRequest(file *models.File) error {
	client := &http.Client{}
	
	req,err := http.NewRequest(http.MethodGet,file.Link,nil)
	if err!=nil{
		return err
	}
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64)",
	)

	res,err := client.Do(req)
	if err!=nil{
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("ERROR :Bad Gateway!")
	}

	disposition := res.Header.Get("Content-Disposition")
	// partialDownSupport := res.Header.Get("Accept-Ranges")

	_ ,parsedDisposition, err  :=  mime.ParseMediaType(disposition)
	if err!= nil{
		return err
	}
	
	var fileName string = "unknown.bin"

	fileName = parsedDisposition["filename"] 
	downloadPath  := helpers.GetDownloadPath()

	fileName = helpers.GetUniqueFilePath(downloadPath,fileName)

	downloadFilePath := filepath.Join(downloadPath,fileName)

	file.Mu.Lock()
	file.Name = fileName
	file.Mu.Unlock()

	FileLeninBytes,err := strconv.Atoi(res.Header.Get("Content-Length"))
	if err!=nil{
		return err
	}
	
	file.Mu.Lock()
	file.TotalSize = FileLeninBytes
	file.Mu.Unlock()

	out, err := os.Create(downloadFilePath)
	if err!=nil{
		return err
	}
	defer out.Close()
	
	file.Mu.Lock()
	file.Path = downloadFilePath
	file.Mu.Unlock()
	buf := make([]byte, 64*1024)


	LastTime := time.Now().UnixMilli()

	lastWriteByte := 0
	currentWriteByte := 0 

	for {
		n, err := res.Body.Read(buf)
		if n > 0 {
			out.Write(buf[:n])
			currentWriteByte+=n
		}
		if err != nil {
			break
		}
		if (time.Now().UnixMilli()-LastTime)> 1000{
			
			timeNow := time.Now().UnixMilli()
			elapsed := float64(timeNow-LastTime) / 1000


			downloadSpeed := int(float64(currentWriteByte-lastWriteByte) / elapsed)
			

			var ETA int
			if downloadSpeed > 0 {
				ETA = (FileLeninBytes - currentWriteByte) / downloadSpeed
			}

			LastTime = timeNow
			lastWriteByte = currentWriteByte


			file.Mu.Lock()
			file.ETA = helpers.CalculateTime(ETA)
			file.Downloaded = currentWriteByte
			file.Mu.Unlock()
		}

	}
	file.Mu.Lock()
	file.IsCompleted = true
	file.Downloaded = currentWriteByte
	file.ETA = "0s"
	file.Mu.Unlock()

	return nil
}