package main

import (
	"blackbox/internal/torrents"
	"blackbox/internal/torrents/helpers"
	"bytes"
	"fmt"
	"os"
)

func main() {
	//TESTING NOWWWWW
	raw, _ := os.ReadFile("HELLO.torrent")
	torrentFull,err := torrents.OpenTorrent(bytes.NewBuffer(raw))
	if err!=nil{
		fmt.Println("Torrent is corrupted")
		os.Exit(1)
	}
	torrentINFO,_ := torrentFull.ToTorrentFile()
	peerID,_ := helpers.GeneratePeerID()
	fmt.Printf("%s\n",peerID)
	url ,_ := torrentINFO.BuildURL([20]byte(peerID),6881)
	fmt.Println(url)
}