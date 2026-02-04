// Package tui - tea.Model (Init, Update, View)
package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/koha90/tui_bots_manager/internal/bot"
)

type Model struct {
	Bots        []bot.Bot
	Cursor      int
	Manager     *bot.Manager
	ConfirmStop string
}

func New(mgr *bot.Manager) Model {
	return Model{
		Manager: mgr,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		func() tea.Msg {
			return BotsLoadedMsg{Bots: m.Manager.List()}
		},
		ListenBotEvents(m.Manager),
	)
}

// Update - обработка и привязка клавиш.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		b := m.Bots[m.Cursor]
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
			if b.Status() == bot.Stopped {
				return m, StartBotCmd(b, m.Manager)
			}
		case "x", "ч":
			currentStatus := m.Manager.Status(b.ID())
			if m.ConfirmStop != b.ID() {
				m.ConfirmStop = b.ID()
				return m, nil
			}
			m.ConfirmStop = ""
			if currentStatus == bot.Running {
				return m, StopBotCmd(b, m.Manager)
			}
		}

	case BotsLoadedMsg:
		m.Bots = msg.Bots
		m.Cursor = 0

	case TickMsg:
		return m, TickCmd()

	case BotStateChangeMsg:
		return m, ListenBotEvents(m.Manager)
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

		s += fmt.Sprintf("%s %-20s %s\n", cursor, b.ID(), style.Render(status.String()))
	}

	// Предупреждение об остановке бота.
	if m.ConfirmStop != "" && m.Manager.Status(m.ConfirmStop) == bot.Running {
		s += redSoft.Render(ConfirmStopMsg)
	}

	// сообщение об ошибке, если есть
	for _, b := range m.Bots {
		if m.Manager.Status(b.ID()) == bot.Error {
			s += "\n" + redSoft.Render(ErrorBotMsg)
		}
	}
	s += "\n[s] start  [x] stop  [q] выход\n"
	return s
}
