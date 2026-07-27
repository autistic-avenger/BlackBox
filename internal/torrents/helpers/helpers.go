package helpers

import (
	"crypto/rand"
	"fmt"
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
	peerID := []byte("-TR4130-")
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
	_, err = rand.Read(randomBytes)	
	if err!=nil{
		return nil,err
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