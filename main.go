package main

import (
	"blackbox/assets"
	"log"
	"time"

	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var Logo string = assets.GetLogo() 

type File struct{
	path string
	name string 
	progress float32
	isSelected bool
	isCompleted bool
}

func initModel(m *model){
	m.isLoading = true
}


type model struct{
	height int
	width int
	AddLink textarea.Model
	Files []File
	isLoading bool
}

type IOFetchMsg struct{}

func fetchFiles() tea.Msg {
	time.Sleep(5* time.Second)
	return IOFetchMsg{}
}

func (m model) Init() tea.Cmd { 
	return fetchFiles
}

func (m model) Update(msg tea.Msg) (tea.Model,tea.Cmd){
	switch msg := msg.(type){
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width  = msg.Width
	case tea.KeyMsg:
		switch msg.String(){
		case "q":
			return m, tea.Quit
		}
	}
	
	return m,nil
}

func (m model) View() tea.View{

	v:= tea.NewView(lipgloss.Place(m.width,7,lipgloss.Center,lipgloss.Center,Logo))

	v.WindowTitle = "BlackBox"
	v.BackgroundColor = lipgloss.Color("#282836")
	v.AltScreen = true
	return v
}

func main() {
	m := model{}
	initModel(&m)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err!=nil{
		log.Fatal("Error Starting tea program | Crashed...")
	}

}