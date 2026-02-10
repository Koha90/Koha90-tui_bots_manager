// Package tui - lipgloss цвета и стили.
package tui

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/koha90/tui_bots_manager/internal/bot"
)

var (
	greenSoft = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7DCEA0"))

	redSoft = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E74C3c"))

	graySoft = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAB7B8"))

	yellowSoft = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F7DC49"))
	cursorSoft = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F7DC6F"))
)

var (
	headerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#45ada8"))

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#c7f464"))

	footerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#45ada8"))
)

func statusStyle(s bot.Status) lipgloss.Style {
	switch s {
	case bot.Running:
		return greenSoft
	case bot.Starting:
		return yellowSoft
	case bot.Error:
		return redSoft
	default:
		return graySoft
	}
}

func botStyle(b bot.Bot) lipgloss.Style {
	switch b.Status() {
	case bot.Running:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	case bot.Error:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	default:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	}
}
