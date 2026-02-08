package tui

type Layout struct {
	Cols int
	Rows int
}

func calcLayout(total int, maxCols int) Layout {
	cols := maxCols
	if total < cols {
		cols = 1
	}

	rows := (total + cols - 1) / cols

	return Layout{
		Cols: cols,
		Rows: rows,
	}
}

// usableHeight - учёт высоты экрана.
func usableHeight(total int) int {
	const reserved = 6 // заголовок + футер + сообщения
	return total - reserved
}

// // buildGroupedView - рендер групп с разрывами.
// func buildGroupedView(
// 	bots []bot.Bot,
// 	cursor CurorPos,
// 	spinner string,
// 	width int,
// 	height int,
// ) string {
// 	if width == 0 || height == 0 {
// 		return buildGridView(bots, cursor, spinner, 9999)
// 	}
//
// 	rowsAvailable := usableHeight(height)
//
// 	var out strings.Builder
// 	linesUsed := 0
//
// 	for gi, group := range groups {
// 		groupCursor := gi * 10
// 		if groupCursor < 0 || groupCursor >= len(group) {
// 			groupCursor = -1
// 		}
//
// 		view := buildGridView(group, groupCursor, spinner, width)
// 		lines := strings.Count(view, "\n")
//
// 		// если не влезаем - визульный разрыв
// 		if gi > 0 && linesUsed+lines > rowsAvailable {
// 			break
// 		}
//
// 		out.WriteString(view)
// 		linesUsed += lines
// 	}
//
// 	return out.String()
