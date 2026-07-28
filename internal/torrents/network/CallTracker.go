package network

import (
	"blackbox/internal/torrents"
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/jackpal/bencode-go"
)

func CallTracker(url string) (*torrents.AnnouncerResponse , error) {
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
	
	var responseStruct torrents.AnnouncerResponse
	err = bencode.Unmarshal(bytes.NewBuffer(responseBytes),&responseStruct)
	if err!=nil{
		return nil,err
	}
	
	return &responseStruct,nil
}