package tui

type Grid struct {
	Cols int
	Rows int // максимальное число строк в колонке
}

func calcGrid(total, width int) Grid {
	if total == 0 {
		return Grid{
			Cols: 1,
			Rows: 0,
		}
	}

	cols := max(1, width/ColWidth)
	rows := (total + cols - 1) / cols

	return Grid{
		Cols: cols,
		Rows: rows,
	}
}
