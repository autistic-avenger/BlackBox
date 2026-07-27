package helpers

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
)

func PieceHash(pieces string) [][20]byte {

	hashes := make([][20]byte, len(pieces)/20)

	for i := range hashes {
		copy(hashes[i][:], pieces[20*i:20*(i+1)])
	}
	return hashes
}

func GeneratePeerID() ([]byte,error) {
	CLIENT_CODE := "-B69420-"

	peerID := []byte(CLIENT_CODE)
	randomBytes := make([]byte,12)

	Appdata := os.Getenv("LOCALAPPDATA")
	if Appdata == ""{
		return nil, fmt.Errorf("Appdata not Found!")
	}

	blackboxPath := filepath.Join(Appdata,"blackbox")
	
	idFile := filepath.Join(blackboxPath,"peerID.bin")
	_, err := os.Stat(idFile)
	if err == nil{
		peerID, err = os.ReadFile(idFile)
		if err!=nil{
			return nil,err
		}
		return peerID,nil
	} 

	abcs := "abcdefghijklmnopqrstuvwxyz0123456789"
	for i := range randomBytes {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(abcs))))
		if err != nil {
			return nil, err
		}
		randomBytes[i] = abcs[n.Int64()]
	}


	peerID = append(peerID, randomBytes...)
	err = os.MkdirAll(blackboxPath,0755)
	if err!=nil{
		return nil, err 
	}

	err = os.WriteFile(idFile,peerID,0644)
	if err!=nil{
		return nil, err 
	}

	return peerID, nil

}