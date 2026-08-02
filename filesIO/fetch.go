package filesio

import (
	"encoding/csv"
	"os"

	tea "charm.land/bubbletea/v2"
)

type IOFetchMsg struct{
	files [][]string
}

func FetchFiles() tea.Msg {
	filePath := GetPath()
	file, err := os.Open(filePath)
	if err!=nil{
		return IOFetchMsg{
			files: [][]string{},
		}
	}
	reader := csv.NewReader(file)

	records,err := reader.ReadAll()
	if err!=nil{
		return IOFetchMsg{
			files: [][]string{},
		}
	}
	return IOFetchMsg{
		files: records,
	}
}
