package tui

import (
	"fmt"
	"strings"

	"github.com/koha90/tui_bots_manager/internal/bot"
)

const (
	BotsPerColumn = 10
	ColWidth      = 30
)

const (
	idWidth     = 20 // ширина для ID
	statusWidth = 8  // ширина для статуса
)

type CurorPos struct {
	Col int
	Row int
}

func buildGridView(
	bots []bot.Bot,
	cursor CurorPos,
	spinner string,
	width int,
) string {
	if len(bots) == 0 {
		return ""
	}

	cols := calcColumns(width)
	var out strings.Builder

	for row := range BotsPerColumn {
		for col := range cols {
			i := col*BotsPerColumn + row
			if i >= len(bots) {
				continue
			}

			b := bots[i]

			cursorChar := " "
			if cursor.Row == row && cursor.Col == col {
				cursorChar = cursorSoft.Render("")
			}

			spin := "  "
			if b.Status() == bot.Starting {
				spin = spinner
				if len(spin) < 2 {
					spin += strings.Repeat("", 2-len(spin))
				}
			}

			style := statusStile(b.Status())

			fmt.Fprintf(
				&out,
				" %s%s%-*s %-*s",
				cursorChar,
				spin,
				idWidth, b.ID(),
				statusWidth, style.Render(b.Status().String()), // statusWidth,
			)
		}
		out.WriteByte('\n')
	}

	return out.String()
}

func calcColumns(width int) int {
	cols := width / ColWidth
	if cols < 1 {
		return 1
	}
	return cols
}
