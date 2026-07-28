package networks

import (
	"blackbox/downloader/helpers"
	"fmt"
	"mime"
	"strconv"
	"time"

	// "strconv"

	"net/http"
	"os"
)

func MakeRequest(url string) error {
	client := &http.Client{}
	
	req,err := http.NewRequest(http.MethodGet,url,nil)
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
	_ ,parsedDisposition, err  :=  mime.ParseMediaType(disposition)
	if err!= nil{
		return err
	}
	
	var fileName string = "unknown.bin"

	fileName = parsedDisposition["filename"] 
	fmt.Printf("\n\nDownloading | %s\n",fileName)

	FileLeninBytes,err := strconv.Atoi(res.Header.Get("Content-Length"))
	if err!=nil{
		return err
	}
	
	fmt.Printf("Size 	    | %s\n",helpers.CalculateSize(FileLeninBytes))

	out, err := os.Create(fileName)
	if err!=nil{
		return err
	}
	defer out.Close()
	
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
		if (time.Now().UnixMilli()-LastTime)>1000{
			
			timeNow := time.Now().UnixMilli()
			elapsed := int(timeNow-LastTime) / 1000

			downloadSpeed := (currentWriteByte-lastWriteByte) / elapsed
			
			ETA 	:=  (FileLeninBytes-currentWriteByte)/downloadSpeed
			LastTime = timeNow
			lastWriteByte = currentWriteByte
			fmt.Printf("\r ETA: %s Speed: %s/s\n",helpers.CalculateTime(ETA),helpers.CalculateSize(downloadSpeed))
		}

	}

	return nil
}