package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct{

}

func (m model) Init() tea.Cmd{
	return nil
}

func (m model) View() string {
	return "TESTING 123\nPress q to quit"
}

func (m model) Update(msg tea.Msg) (tea.Model ,tea.Cmd){
	switch msg := msg.(type){
	case tea.KeyMsg:
		switch msg.String(){
		case "q":
			return m, tea.Quit
		}
	}
	return m ,nil
}

func main(){

	p := tea.NewProgram(model{})
	if _, teaErr := p.Run(); teaErr!=nil{
		log.Panic(teaErr)
	}

}