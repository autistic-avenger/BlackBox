package network

import (
	"io"
	"net/http"
	"os"
)

func CallTracker(url string) ([]byte , error) {
	Client := &http.Client{}
	
	req ,err := http.NewRequest(http.MethodGet,url,nil)
	if err!= nil {
		return nil,err
	}

	res ,err := Client.Do(req)
	if err!=nil{
		return nil,err
	}
	defer res.Body.Close()

	responseBytes ,err := io.ReadAll(res.Body)
	if err!=nil{
		return nil,err
	}
	os.WriteFile("response.torrent",responseBytes,0644)

	return responseBytes,nil
}