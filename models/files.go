package models

import "sync"

type File struct{
	Mu sync.RWMutex
	Link string
	Path string
	Name string 
	Downloaded int
	TotalSize int
	Speed string
	IsSelected bool
	IsCompleted bool
	ETA string
	Error string
	Timeout bool
}