package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/koha90/tui_bots_manager/internal/bot"
)

func loadBotsCmd(mgr *bot.Manager) func() tea.Msg {
	return func() tea.Msg {
		bots := mgr.List()

		result := make([]BotView, 0, len(bots))
		for _, b := range bots {
			result = append(result, BotView{
				ID:     b.ID,
				Status: mapStatus(b.Status),
			})
		}

		return BotsLoadedMsg{Bots: result}
	}
}

func mapStatus(s bot.Status) BotStatus {
	switch s {
	case bot.Running:
		return Running
	case bot.Error:
		return Error
	default:
		return Stopped
	}
}
