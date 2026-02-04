package main

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/koha90/tui_bots_manager/internal/bot"
	"github.com/koha90/tui_bots_manager/internal/tui"
)

func main() {
	mgr := bot.NewManager()

	alpha := bot.NewFake("alpha")
	beta := bot.NewFake("beta")
	gamma := bot.NewFake("gamma")

	mgr.Add(alpha)
	mgr.Add(beta)
	mgr.Add(gamma)

	go func() {
		time.Sleep(2 * time.Second)
		beta.SimulateError()
	}()

	p := tea.NewProgram(
		tui.New(mgr),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
