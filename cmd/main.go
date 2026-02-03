package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/koha90/tui_bots_manager/internal/bot"
	"github.com/koha90/tui_bots_manager/internal/tui"
)

func main() {
	mgr := bot.NewManager()
	p := tea.NewProgram(
		tui.New(mgr),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
