package tui

import (
	"strings"

	"github.com/koha90/tui_bots_manager/internal/bot"
)

// View - отрисовка.
func (m Model) View() string {
	var out strings.Builder

	// --- Header ---
	out.WriteString("👾 Bots List:\n\n")

	// --- Bots Grid ---
	if len(m.Bots) == 0 {
		out.WriteString("(нет ботов)\n")
	} else {
		out.WriteString(
			buildGridView(
				m.Bots,
				m.Cursor,
				m.spinner.View(),
				m.Width,
			),
		)
	}

	out.WriteByte('\n')
	if m.ConfirmStop != "" && m.Manager.Status(m.ConfirmStop) == bot.Running {
		out.WriteString(redSoft.Render(ConfirmStopMsg))
	}
	out.WriteByte('\n')

	hasError := false
	for _, b := range m.Bots {
		if m.Manager.Status(b.ID()) == bot.Error {
			hasError = true
			break
		}
	}
	if hasError {
		out.WriteString(redSoft.Render(ErrorBotMsg))
	}
	// out.WriteByte('\n')

	out.WriteString("\n\n[s] start  [x] stop  [q] выход\n")
	return out.String()
}
