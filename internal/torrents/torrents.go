package torrents

import (
	"blackbox/internal/torrents/helpers"
	"bytes"
	"crypto/sha1"
	"io"

	"github.com/jackpal/bencode-go"
)

type TorrentRaw struct{
	Announce	string 		`bencode:"announce"`
	Comment 	string 		`bencode:"comment"`
	CreatedDate	int    		`bencode:"creation date"`
	Info 		TorrentInfo `bencode:"info"`
}


type TorrentInfo struct{
	Length 		int 		`bencode:"length"`
	Name 		string		`bencode:"name"`
	PieceLen	int 		`bencode:"piece length"`
	Pieces 		string 		`bencode:"pieces"`
}


func OpenTorrent(r io.Reader) (*TorrentRaw,error){
	torrent := TorrentRaw{}
	err := bencode.Unmarshal(r,&torrent)
	if err!=nil{
		return nil,err
	}
	return &torrent, nil 	
}

type TorrentFile struct {
    Announce    string
    InfoHash    [20]byte
    PieceHashes [][20]byte
    PieceLength int
    Length      int
    Name        string
}

func (parsedTorrent TorrentRaw) toTorrentFile() (TorrentFile,error){
	var infoBuf bytes.Buffer
	err := bencode.Marshal(&infoBuf,parsedTorrent.Info)
	if err!=nil{
		return TorrentFile{},err
	}
	infoHash := sha1.Sum(infoBuf.Bytes())
	pieceHash := helpers.PieceHash(parsedTorrent.Info.Pieces)

	tf:= TorrentFile{
		Announce: parsedTorrent.Announce,
		InfoHash: infoHash,
		PieceHashes: pieceHash,
		PieceLength: parsedTorrent.Info.PieceLen,
		Length: parsedTorrent.Info.Length,
		Name: parsedTorrent.Info.Name,
	}
	return tf,nil
}