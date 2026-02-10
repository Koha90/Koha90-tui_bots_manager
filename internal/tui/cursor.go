package tui

type CurorPos struct {
	Col int
	Row int
}

func (n *Navigator) Index(grid Grid) int {
	return n.Cursor.Col*grid.Rows + n.Cursor.Row
}
