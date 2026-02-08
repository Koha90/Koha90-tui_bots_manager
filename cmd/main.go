package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/koha90/tui_bots_manager/internal/bot"
	"github.com/koha90/tui_bots_manager/internal/tui"
)

func main() {
	bots := []bot.Bot{
		bot.NewFake("alpha1"),
		bot.NewFake("alpha2"),
		bot.NewFake("alpha3"),
		bot.NewFake("alpha4"),
		bot.NewFake("alpha5"),
		bot.NewFake("alpha6"),
		bot.NewFake("alpha7"),
		bot.NewFake("alpha8"),
		bot.NewFake("alpha9"),
		bot.NewFake("alpha10"),
		bot.NewFake("alpha11"),
		bot.NewFake("alpha12"),
		bot.NewFake("alpha13"),
		bot.NewFake("alpha14"),
		bot.NewFake("alpha15"),
		bot.NewFake("alpha16"),
		bot.NewFake("alpha17"),
		bot.NewFake("alpha18"),
		bot.NewFake("alpha19"),
		bot.NewFake("alpha20"),
		bot.NewFake("alpha21"),
		bot.NewFake("alpha22"),
		bot.NewFake("alpha23"),
		bot.NewFake("alpha24"),
		bot.NewFake("alpha25"),
		bot.NewFake("alpha26"),
		bot.NewFake("alpha27"),
		bot.NewFake("alpha28"),
		bot.NewFake("alpha29"),
		bot.NewFake("alpha30"),
		bot.NewFake("alpha31"),
		bot.NewFake("alpha32"),
		bot.NewFake("alpha33"),
		bot.NewFake("alpha34"),
		bot.NewFake("alpha35"),
		bot.NewFake("alpha36"),
		bot.NewFake("alpha37"),
		bot.NewFake("alpha38"),
		bot.NewFake("alpha39"),
	}

	mgr := bot.NewManager()

	for _, b := range bots {
		mgr.Add(b)
	}

	p := tea.NewProgram(
		tui.New(mgr),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
