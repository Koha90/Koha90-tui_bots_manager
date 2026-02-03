// Package tui - lipgloss цвета и стили.
package tui

import "github.com/charmbracelet/lipgloss"

var (
	greenSoft = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7DCEA0"))

	redSoft = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E74C3c"))

	graySoft = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAB7B8"))

	yellowSoft = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F7DC49"))
	cursorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F7DC6F"))
)
