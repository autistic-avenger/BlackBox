package main

import (
	"blackbox/downloader"
	"fmt"
	"os"
)

func main() {
	var url string

	fmt.Print("Enter URL: ")
	fmt.Scanln(&url)
	fmt.Printf("\rDownloading...\n")
	err := downloader.DownloadFile(url)
	if err!= nil{
		os.Exit(1)
	}
}