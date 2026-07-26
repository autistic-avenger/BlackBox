package main

import (
	"blackbox/internal/torrents"
	"bytes"
	"fmt"
	"os"
)

func main() {
	//TESTING NOWWWWW
	raw, _ := os.ReadFile("HELLO.torrent")
	torrentFull,_ := torrents.OpenTorrent(bytes.NewBuffer(raw))
	torrentINFO,_ := torrentFull.ToTorrentFile()
	peerID := []byte("PeerIDPeerIDPeerID22")
	url ,_ := torrentINFO.BuildURL([20]byte(peerID),6881)
	fmt.Println(url)
}