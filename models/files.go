package models

import "sync"

type File struct{
	Mu sync.RWMutex
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