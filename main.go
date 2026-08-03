package main

import (
	"blackbox/assets"
	"blackbox/downloader"
	"blackbox/downloader/helpers"
	filesio "blackbox/filesIO"
	"blackbox/models"
	"blackbox/ticker"
	"log"
	"slices"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var Logo string = assets.GetLogo() 

var (
	InputStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#6B61BA"))
	HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#B7B2E6"))
	NotFoundStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#85869E"))
	DivStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#FF9A9B"))
)


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
	Files []*models.File
	isLoading bool
}


func (m model) Init() tea.Cmd { 
	return tea.Batch(
		filesio.FetchFiles,
		ticker.TickTimer(),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model,tea.Cmd){
	var cmd tea.Cmd

	switch msg := msg.(type){
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width  = msg.Width
		m.AddLink.SetWidth(min(60,m.width-3))
		return m, nil
	case ticker.TickMsg:
		return m,ticker.TickTimer()
	case filesio.IOFetchMsg:
		m.isLoading = false
		return m,nil
	case tea.KeyMsg:
		switch msg.String(){
		case "q":
			return m, tea.Quit
		case "/":
			cmd = m.AddLink.Focus()
			m.AddLink.Placeholder = ""
			return m,cmd
		case "enter":
			if m.AddLink.Focused(){
				var File models.File
				File.Link = m.AddLink.Value()
				downloader.DownloadFile(&File)
				m.Files = append(m.Files, &File)
				m.AddLink.Blur()
				m.AddLink.SetValue("")
				m.AddLink.Placeholder = `Press "/" to add link`

			}
		case "esc":
			if m.AddLink.Focused(){
				m.AddLink.Blur()
				m.AddLink.Placeholder = `Press "/" to add link`
			}
		}
	}
	m.AddLink, cmd = m.AddLink.Update(msg)
	return m, cmd
}

func (m model) View() tea.View{
	theLogo := lipgloss.Place(
		m.width,
		7,
		lipgloss.Center,
		lipgloss.Center,
		Logo,
	)

	inputTag := InputStyle.Render(m.AddLink.View())

	var Downloads []string

	for _, fileData := range m.Files{
		FileStyle := DivStyle.Render(
			lipgloss.Place(
				min(68,m.width-2),
				5,
				lipgloss.Center,
				lipgloss.Top,
				lipgloss.JoinVertical(
					lipgloss.Top,
					lipgloss.JoinHorizontal(
						lipgloss.Center,
						lipgloss.NewStyle().MaxWidth(34).Foreground(lipgloss.Color("#ffffff")).Bold(true).Render(lipgloss.Place(
							34,
							1,
							lipgloss.Left,
							lipgloss.Center,
							" "+helpers.Truncate(fileData.Name,33),
						)),
						lipgloss.NewStyle().MaxWidth(34).Foreground(lipgloss.Color("#ffffff")).Bold(true).Render(lipgloss.Place(
							34,
							1,
							lipgloss.Right,
							lipgloss.Center,
							"▶  ",
						)),
					),
					"",
					lipgloss.Place(
						min(68,m.width-2),
						2,
						lipgloss.Left,
						lipgloss.Center,
						lipgloss.JoinHorizontal(
							lipgloss.Left,
							" ",
							lipgloss.NewStyle().Background(lipgloss.Color("#67658d")).Render(lipgloss.Place(
								57,
								1,
								lipgloss.Left,
								lipgloss.Center,
								"",
							)),
							"      97%",
						),

					),
				),
			),
		)
		Downloads = append(Downloads, FileStyle)
	}

	slices.Reverse(Downloads)
	
	var DownloadSection string
	if len(m.Files) == 0 {
		DownloadSection = lipgloss.Place(
			m.width,
			m.height/2,
			lipgloss.Center,
			lipgloss.Center,
			NotFoundStyle.Render("Download list is Empty."))
	}else{
		DownloadSection = lipgloss.JoinVertical(
			lipgloss.Top,
			Downloads...
		)
	}

	body := lipgloss.Place(
		m.width,
		m.height-2,
		lipgloss.Center,
		lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Center,
			theLogo,
			inputTag,
			DownloadSection,
		),
	)	

	helpTxt := "/: Add link • ↑↓: Select • Esc: back • q: Quit"
	helpTag := lipgloss.Place(
		m.width,
		1,
		lipgloss.Center,
		lipgloss.Bottom,
		HelpStyle.Render(helpTxt),
	)
	
	var v tea.View
	if m.isLoading{
		v = tea.NewView(lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			"Loading..."),
		)
	}else{
		v= tea.NewView(lipgloss.JoinVertical(
			lipgloss.Center,
			body,
			helpTag),
		)
	}
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