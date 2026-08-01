package main

import (
	"blackbox/assets"
	"log"
	"time"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var Logo string = assets.GetLogo() 

var (
	InputStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#6B61BA"))
)

type File struct{
	path string
	name string 
	progress float32
	isSelected bool
	isCompleted bool
}

func initModel(m *model){
	ti := textinput.New()
	ti.Placeholder = `Press "/" to add link.`
	ti.Prompt = " >  "

	m.AddLink = ti
	m.isLoading = true
}


type model struct{
	height int
	width int
	AddLink textinput.Model
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
	var cmd tea.Cmd

	switch msg := msg.(type){
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width  = msg.Width
		m.AddLink.SetWidth(min(60,m.width-3))
		return m, nil
	case tea.KeyMsg:
		switch msg.String(){
		case "q":
			return m, tea.Quit
		case "/":
			cmd = m.AddLink.Focus()
			return m,cmd
		case "enter":
			if m.AddLink.Focused(){
				m.AddLink.Blur()
				m.AddLink.SetValue("")
			}
		case "esc":
			if m.AddLink.Focused(){
				m.AddLink.Blur()
			}	
		}
	}
	m.AddLink, cmd = m.AddLink.Update(msg)
	return m, cmd
}

func (m model) View() tea.View{
	theLogo := lipgloss.Place(m.width,7,lipgloss.Center,lipgloss.Center,Logo)

	v:= tea.NewView(lipgloss.JoinVertical(lipgloss.Center,theLogo, InputStyle.Render(m.AddLink.View())))
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