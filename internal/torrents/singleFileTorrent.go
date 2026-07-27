package torrents

import (
	"blackbox/internal/torrents/helpers"
	"bytes"
	"crypto/sha1"
	"io"
	"net/url"
	"strconv"

	"github.com/jackpal/bencode-go"
)

type TorrentRaw struct{
	Announce	string 		`bencode:"announce"`
	AnnounceList [][]string `bencode:"announce-list"`
	Comment 	string 		`bencode:"comment"`
	CreatedDate	int    		`bencode:"creation date"`
	Info 		Info 		`bencode:"info"`
}


type  Info struct{
	Length 		int 			`bencode:"length,omitempty"`
	Name 		string			`bencode:"name"`
	PieceLen	int 			`bencode:"piece length"`
	Pieces 		string 			`bencode:"pieces"`
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

func (parsedTorrent TorrentRaw) ToTorrentFile() (TorrentFile,error){
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
		Name: parsedTorrent.Info.Name,
	}
	return tf,nil
}



func (tf TorrentFile) BuildURL(peerID [20]byte ,port int) (string, error) {
	baseURL, err := url.Parse(tf.Announce)
	if err != nil {
		return "", err
	}
	
	querry := baseURL.Query()
	querry.Add("info_hash",string(tf.InfoHash[:]))
	querry.Add("peer_id",string(peerID[:]))
	querry.Add("port",strconv.Itoa(port))
	querry.Add("uploaded","0")
	querry.Add("downloaded","0")
	querry.Add("compact","1")
	querry.Add("left",strconv.Itoa(tf.Length))

	baseURL.RawQuery = querry.Encode()	
	return baseURL.String(), nil
}