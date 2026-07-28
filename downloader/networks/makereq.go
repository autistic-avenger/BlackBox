package networks

import (
	"fmt"
	"mime"

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

	disposition := res.Header.Get("Content-Disposition")
	mediaType ,parsedDisposition, err  :=  mime.ParseMediaType(disposition)
	if err!= nil{
		return err
	}
	fmt.Println("MeidaType :",mediaType)
	fmt.Println("parsed :",parsedDisposition)

	fileName := parsedDisposition["filename"]

	out, err := os.Create(fileName)
	if err!=nil{
		return err
	}
	defer out.Close()
	
	buf := make([]byte, 32*1024)

	for {
    n, err := res.Body.Read(buf)
    if n > 0 {
        out.Write(buf[:n])
    }
    if err != nil {
        break
    }
	}
		
	return nil
}