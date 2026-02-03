// Package tui - tea.Model (Init, Update, View)
package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/koha90/tui_bots_manager/internal/bot"
)

type Model struct {
	Bots    []bot.Bot
	Cursor  int
	Manager *bot.Manager
}

func New(mgr *bot.Manager) Model {
	return Model{
		Manager: mgr,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		TickCmd(),
		func() tea.Msg {
			return BotsLoadedMsg{
				Bots: m.Manager.List(),
			}
		},
	)
}

// Update - обработка и привязка клавиш.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "й":
			return m, tea.Quit
		case "up", "k", "л":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j", "о":
			if m.Cursor < len(m.Bots)-1 {
				m.Cursor++
			}
		case "s", "ы":
			bot := m.Bots[m.Cursor]
			return m, StartBotCmd(bot)
		case "x", "ч":
			bot := m.Bots[m.Cursor]
			return m, StopBotCmd(bot)
		}

	case BotsLoadedMsg:
		m.Bots = msg.Bots
		m.Cursor = 0

	case TickMsg:
		return m, TickCmd()
	}

	return m, nil
}

// View - отрисовка.
func (m Model) View() string {
	s := "👾 Bots:\n\n"

	for i, b := range m.Bots {
		cursor := " "
		if m.Cursor == i {
			cursor = cursorStyle.Render("")
		}

		status := b.Status()
		style := graySoft

		switch status {
		case bot.Running:
			style = greenSoft
		case bot.Starting:
			style = yellowSoft
		case bot.Stopped:
			style = graySoft
		case bot.Error:
			style = redSoft
		}

		s += fmt.Sprintf(
			"%s %-20s %s\n",
			cursor, b.ID(),
			style.Render(status.String()),
		)
	}

	s += "\n[s] start  [x] stop  [q] выход\n"
	return s
}
