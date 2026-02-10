package tui

// View - отрисовка.
func (m Model) View() string {
	if m.Width == 0 || m.Height == 0 {
		return "loading..."
	}

	return Render(m.Bots, m.Nav.Cursor, m.spinner.View(), m.Grid, m.Height)
}
