package ticker

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

type TickMsg time.Time


func TickTimer() tea.Cmd {
	return tea.Tick(500*time.Millisecond,func(t time.Time) tea.Msg {
		return  TickMsg(t)
	})
}