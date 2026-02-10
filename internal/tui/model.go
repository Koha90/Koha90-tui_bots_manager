// Package tui - tea.Model (Init, Update, View)
package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/koha90/tui_bots_manager/internal/bot"
)

type Model struct {
	Bots        []bot.Bot
	Nav         Navigator
	Grid        Grid
	Manager     *bot.Manager
	ConfirmStop string
	spinner     spinner.Model

	Width  int
	Height int
}

func New(mgr *bot.Manager) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return Model{
		Manager: mgr,
		spinner: s,
		Nav:     NewNavigator(),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		func() tea.Msg {
			return BotsLoadedMsg{Bots: m.Manager.List()}
		},
		ListenBotEvents(m.Manager),
	)
}
