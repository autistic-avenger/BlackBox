package main

import (
	"blackbox/downloader"
	"fmt"
)

func main() {
	var url string

	fmt.Printf("Enter URL: ")
	fmt.Scanln(&url)
	err := downloader.DownloadFile(url)
	if err!= nil{
		fmt.Println(err)
	}
}