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
	Cursor      CurorPos
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

// Update - обработка и привязка клавиш.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var spinnerCmd tea.Cmd

	m.spinner, spinnerCmd = m.spinner.Update(msg)
	cmds = append(cmds, spinnerCmd)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

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
			if m.currentIndex() == 0 {
				last := len(m.Bots) - 1
				m.Cursor.Col = last / BotsPerColumn
				m.Cursor.Row = last % BotsPerColumn
				m.ConfirmStop = ""
				return m, nil
			}
			if m.Cursor.Row > 0 {
				m.Cursor.Row--
			} else {
				m.Cursor.Col--
				m.Cursor.Row = BotsPerColumn - 1
			}
			m.normalizeCursor()
			m.ConfirmStop = ""

		case "down", "j", "о":
			if m.currentIndex() == len(m.Bots)-1 {
				m.Cursor.Col = 0
				m.Cursor.Row = 0
				m.ConfirmStop = ""
				return m, nil
			}

			start := m.Cursor.Col * BotsPerColumn
			remaining := len(m.Bots) - start

			maxRows := BotsPerColumn
			remaining = min(remaining, BotsPerColumn)
			maxRows = remaining

			if m.Cursor.Row+1 < maxRows {
				m.Cursor.Row++
			} else {
				m.Cursor.Col++
				m.Cursor.Row = 0
			}

			m.normalizeCursor()
			m.ConfirmStop = ""

		case "left", "h", "р":
			maxCols := m.maxCols()

			if m.Cursor.Col == 0 {
				m.Cursor.Col = maxCols - 1
			} else {
				m.Cursor.Col--
			}
			m.normalizeCursor()
			m.ConfirmStop = ""

		case "right", "l", "д":
			maxCols := m.maxCols()

			if m.Cursor.Col == maxCols-1 {
				m.Cursor.Col = 0
			} else {
				m.Cursor.Col++
			}
			m.normalizeCursor()
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
		m.Cursor.Col = 0

	case TickMsg:
		cmds = append(cmds, TickCmd())

	case BotStateChangeMsg:
		cmds = append(cmds, ListenBotEvents(m.Manager))
	}

	return m, tea.Batch(cmds...)
}

func (m Model) currentIndex() int {
	return m.Cursor.Col*BotsPerColumn + m.Cursor.Row
}

func (m Model) currentBot() (bot.Bot, bool) {
	i := m.currentIndex()
	if i < 0 || i >= len(m.Bots) {
		return nil, false
	}
	return m.Bots[i], true
}

func (m *Model) normalizeCursor() {
	cols := calcColumns(m.Width)

	if m.Cursor.Col < 0 {
		m.Cursor.Col = 0
	}
	if m.Cursor.Col >= cols {
		m.Cursor.Col = cols - 1
	}

	start := m.Cursor.Col * BotsPerColumn
	remaining := len(m.Bots) - start

	maxRows := BotsPerColumn
	remaining = min(remaining, BotsPerColumn)
	maxRows = remaining

	if m.Cursor.Row < 0 {
		m.Cursor.Row = 0
	}
	if m.Cursor.Row >= maxRows {
		m.Cursor.Row = maxRows - 1
	}
}

func (m Model) maxCols() int {
	cols := calcColumns(m.Width)
	needed := (len(m.Bots) + BotsPerColumn - 1) / BotsPerColumn
	if cols > needed {
		return needed
	}
	return cols
}
