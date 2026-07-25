package helpers


func PieceHash(pieces string) [][20]byte {

	hashes := make([][20]byte, len(pieces)/20)

	for i :=range hashes{
		copy(hashes[i][:],pieces[20*i:20*(i+1)])
	}
	return hashes
}