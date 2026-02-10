package tui

type Navigator struct {
	Cursor CurorPos
}

func NewNavigator() Navigator {
	return Navigator{
		Cursor: CurorPos{Row: 0, Col: 0},
	}
}

// -----------------------------------------------------

func (n *Navigator) normalize(total int, grid Grid) {
	if total == 0 || grid.Cols == 0 || grid.Rows == 0 {
		n.Cursor = CurorPos{}
		return
	}

	// ---- нормализация колонки ----
	if n.Cursor.Col < 0 {
		n.Cursor.Col = 0
	}
	if n.Cursor.Col >= grid.Cols {
		n.Cursor.Col = grid.Cols - 1
	}

	// ---- строки в текущей колонке ----
	start := n.Cursor.Col * grid.Rows
	maxRows := min(grid.Rows, total-start)

	if n.Cursor.Row < 0 {
		n.Cursor.Row = 0
	}
	if n.Cursor.Row >= maxRows {
		n.Cursor.Row = maxRows - 1
	}
}

// ------------------------------------------------------

// Up - method of navigation to up.
func (n *Navigator) Up(total int, grid Grid) {
	n.Cursor.Row--
	if n.Cursor.Row < 0 {
		n.Cursor.Col = (n.Cursor.Col - 1 + grid.Cols) % grid.Cols
		start := n.Cursor.Col * grid.Rows
		n.Cursor.Row = min(grid.Rows, total-start) - 1
	}
	n.normalize(total, grid)
}

// Down - method of navigation to down.
func (n *Navigator) Down(total int, grid Grid) {
	start := n.Cursor.Col * grid.Rows
	maxRows := min(grid.Rows, total-start)

	n.Cursor.Row++
	if n.Cursor.Row >= maxRows {
		n.Cursor.Row = 0
		n.Cursor.Col = (n.Cursor.Col + 1) % grid.Cols
	}
	n.normalize(total, grid)
}

// Left - method of navigation to left.
func (n *Navigator) Left(total int, grid Grid) {
	n.Cursor.Col = (n.Cursor.Col - 1 + grid.Cols) % grid.Cols
	n.normalize(total, grid)
}

// Right - method of navigation to right.
func (n *Navigator) Right(total int, grid Grid) {
	n.Cursor.Col = (n.Cursor.Col + 1) % grid.Cols
	n.normalize(total, grid)
}
