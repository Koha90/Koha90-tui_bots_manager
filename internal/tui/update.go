package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/koha90/tui_bots_manager/internal/bot"
)

// Update - обработка и привязка клавиш.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var spinnerCmd tea.Cmd

	m.spinner, spinnerCmd = m.spinner.Update(msg)
	cmds = append(cmds, spinnerCmd)

	grid := m.Grid

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Grid = calcGrid(len(m.Bots), m.Width)
		return m, nil

	case tea.KeyMsg:
		if len(m.Bots) == 0 {
			break
		}

		b, ok := m.currentBot()
		if !ok {
			return m, nil
		}

		switch msg.String() {
		case "ctrl+c", "q", "й":
			return m, tea.Quit

		case "up", "k", "л":
			m.Nav.Up(len(m.Bots), grid)
			m.ConfirmStop = ""

		case "down", "j", "о":
			m.Nav.Down(len(m.Bots), grid)
			m.ConfirmStop = ""

		case "left", "h", "р":
			m.Nav.Left(len(m.Bots), grid)
			m.ConfirmStop = ""

		case "right", "l", "д":
			m.Nav.Right(len(m.Bots), grid)
			m.ConfirmStop = ""

		case "s", "ы":
			if b.Status() == bot.Stopped {
				cmds = append(cmds, StartBotCmd(b, m.Manager))
			}

		case "x", "ч":
			currentStatus := m.Manager.Status(b.ID())
			if m.ConfirmStop != b.ID() {
				m.ConfirmStop = b.ID()
				return m, nil
			} else {
				m.ConfirmStop = ""
				if currentStatus == bot.Running {
					cmds = append(cmds, StopBotCmd(b, m.Manager))
				}
			}
		}

	case BotsLoadedMsg:
		m.Bots = msg.Bots
		m.Nav.Cursor.Col = 0
		m.Grid = calcGrid(len(m.Bots), m.Width)

	case TickMsg:
		cmds = append(cmds, TickCmd())

	case BotStateChangeMsg:
		cmds = append(cmds, ListenBotEvents(m.Manager))
	}

	return m, tea.Batch(cmds...)
}

func (m Model) currentIndex() int {
	i := m.Nav.Cursor.Col*m.Grid.Rows + m.Nav.Cursor.Row
	if i >= len(m.Bots) {
		i = len(m.Bots) - 1
	}
	if i < 0 {
		i = 0
	}
	return i
}

func (m Model) currentBot() (bot.Bot, bool) {
	i := m.currentIndex()
	if i < 0 || i >= len(m.Bots) {
		return nil, false
	}
	return m.Bots[i], true
}
