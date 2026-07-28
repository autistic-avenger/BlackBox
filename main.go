package main

import (
	"blackbox/internal/torrents"
	"blackbox/internal/torrents/helpers"
	"blackbox/internal/torrents/network"
	"bytes"
	"fmt"
	"net/netip"
	"os"
)

func main() {
	//TESTING NOWWWWW
	raw, _ := os.ReadFile("ubuntu.torrent")
	torrentFull,err := torrents.OpenTorrent(bytes.NewBuffer(raw))
	if err!=nil{
		fmt.Println("Torrent is corrupted..")
		os.Exit(1)
	}
	torrentINFO,_ := torrentFull.ToTorrentFile()
	peerID,_ := helpers.GeneratePeerID()
	url ,_ := torrentINFO.BuildURL([20]byte(peerID),6881)
	response, err := network.CallTracker(url)
	if err!=nil{
		fmt.Println("ERROR SENDING REQ",err)
	}
	ip,_ := netip.ParseAddr(response.Peers[0].IP)
	fmt.Println(ip.String())
	
}