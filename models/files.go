package models


type File struct{
	Link string
	Path string
	Name string 
	Downloaded string
	TotalSize string
	Speed string
	IsSelected bool
	IsCompleted bool
	ETA string
}