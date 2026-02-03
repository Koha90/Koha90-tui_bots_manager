// Package tui - tea.Model (Init, Update, View)
package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/koha90/tui_bots_manager/internal/bot"
)

type BotStatus int

const (
	Stopped BotStatus = iota
	Running
	Error
)

type BotView struct {
	ID     string
	Status BotStatus
}

type Model struct {
	Bots    []BotView
	Cursor  int
	Manager *bot.Manager
}

func New(mgr *bot.Manager) Model {
	return Model{
		Manager: mgr,
	}
}

func (m Model) Init() tea.Cmd {
	return loadBotsCmd(m.Manager)
}

// Update - обработка и привязка клавиш.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Bots)-1 {
				m.Cursor++
			}
		case "s":
			if len(m.Bots) > 0 {
				bot := m.Bots[m.Cursor]
				return m, func() tea.Msg {
					return StartBotMsg{ID: bot.ID}
				}
			}
		case "x":
			if len(m.Bots) > 0 {
				bot := m.Bots[m.Cursor]
				return m, func() tea.Msg {
					return StartBotMsg{ID: bot.ID}
				}
			}
		}

	case BotsLoadedMsg:
		m.Bots = msg.Bots
		m.Cursor = 0

	case StartBotMsg:
		for i := range m.Bots {
			if m.Bots[i].ID == msg.ID {
				m.Bots[i].Status = Running
			}
		}

	case StopBotMsg:
		for i := range m.Bots {
			if m.Bots[i].ID == msg.ID {
				m.Bots[i].Status = Stopped
			}
		}
	}

	return m, nil
}

// View - отрисовка.
func (m Model) View() string {
	s := "👾 Bots:\n\n"

	for i, bot := range m.Bots {
		cursor := " "
		if m.Cursor == i {
			cursor = cursorStyle.Render("")
		}

		status := ""
		style := graySoft

		switch bot.Status {
		case Running:
			status = "● running"
			style = greenSoft
		case Stopped:
			status = "○ stopped"
			style = graySoft
		case Error:
			status = "✖ error"
			style = redSoft
		}

		s += fmt.Sprintf(
			"%s %-20s %s\n",
			cursor, bot.ID,
			style.Render(status),
		)
	}

	s += "\n[s] start  [x] stop  [q] выход\n"
	return s
}
