package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/koha90/tui_bots_manager/internal/bot"
)

// ----------------- HEADER / INFO / FOOTER --------------------------

func renderHeader() string {
	return headerStyle.Render("\n\n👾 Bots Manager / Управление ботами\n\n")
}

func renderInfo(bots []bot.Bot) string {
	var total, running, stopped, failed int

	for _, b := range bots {
		total++
		switch b.Status() {
		case bot.Running:
			running++
		case bot.Stopped:
			stopped++
		case bot.Error:
			failed++
		}
	}

	return infoStyle.Render(
		fmt.Sprintf(
			"total: %d | running: %d | stopped: %d | error: %d",
			total, running, stopped, failed,
		),
	)
}

func renderFooter() string {
	return footerStyle.Render("[s] start  [x] stop  [q]quit")
}

// ------------------ CELL -----------------------

func renderCell(
	b bot.Bot,
	selected bool,
	cursor bool,
	spinner string,
	width int,
) string {
	style := botStyle(b)

	if selected {
		style = style.Background(lipgloss.Color("236"))
	}

	prefix := "  "

	switch b.Status() {
	case bot.Starting:
		prefix = spinner
	case bot.Running:
		prefix = "  "
	case bot.Error:
		prefix = " "
	}
	if cursor {
		prefix = " "
	}

	content := prefix + b.ID()

	return style.
		Width(width).
		MaxWidth(width).
		PaddingRight(1).
		Render(content)
}

// -------------------------- CONTENT -------------------------

func renderContent(
	bots []bot.Bot,
	cursor CurorPos,
	spinner string,
	grid Grid,
) string {
	if len(bots) == 0 {
		return "(Нет ботов)"
	}

	var out strings.Builder

	for r := 0; r < grid.Rows; r++ {
		for c := 0; c < grid.Cols; c++ {
			i := c*grid.Rows + r
			if i >= len(bots) {
				out.WriteString(strings.Repeat(" ", ColWidth))
				continue
			}

			cursorHere := cursor.Row == r && cursor.Col == c

			cell := renderCell(
				bots[i],
				false,
				cursorHere,
				spinner,
				ColWidth,
			)

			out.WriteString(cell)

		}
		out.WriteByte('\n')
	}
	out.WriteByte('\n')

	return out.String()
}

// ------------------- RENDER --------------------------------------------

func Render(
	bots []bot.Bot,
	cursor CurorPos,
	spinner string,
	grid Grid,
	h int,
) string {
	header := renderHeader()
	info := renderInfo(bots)
	footer := renderFooter()

	used := lipgloss.Height(header) +
		lipgloss.Height(info) +
		lipgloss.Height(footer)

	viewportH := max(0, h-used)

	content := renderContent(bots, cursor, spinner, grid)

	pad := max(0, viewportH-lipgloss.Height(content))

	return lipgloss.JoinVertical(
		lipgloss.Left, header, content, strings.Repeat("\n", pad), info, footer)
}
