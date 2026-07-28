package downloader

import "blackbox/downloader/networks"

func DownloadFile(fileUrl string) error {
	err := networks.MakeRequest(fileUrl)
	if err!=nil{
		return err
	}
	return nil
	
}