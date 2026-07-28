package network

import (
	"fmt"
	"net"
	"time"
)

func CallPeer(peerIP string,infohash []byte,peerId []byte) error {
	conn, err := net.DialTimeout("tcp", peerIP, 3*time.Second)
	if err != nil {
		return err
	}
	defer conn.Close()

	handshake := make([]byte,0,68)
	handshake = append(handshake, 19)
	handshake = append(handshake, []byte("BitTorrent protocol")...)
	handshake = append(handshake, make([]byte, 8)...)
	handshake = append(handshake, infohash[:]...)
	handshake = append(handshake, peerId[:]...)

	_, err = conn.Write(handshake)
	if err != nil {
		return err
	}

	response := make([]byte, 68)

	_, err = conn.Read(response)
	if err != nil {
		return err
	}

	fmt.Printf("Peer handshake: %+v\n", response)

	return nil
}