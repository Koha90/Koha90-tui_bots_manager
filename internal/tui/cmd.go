package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/koha90/tui_bots_manager/internal/bot"
)

type BotStateChangeMsg struct {
	ID     string
	Status bot.Status
}

func StartBotCmd(b bot.Bot, mgr *bot.Manager) tea.Cmd {
	return func() tea.Msg {
		_ = mgr.Start(b.ID())
		return nil
	}
}

func StopBotCmd(b bot.Bot, mgr *bot.Manager) tea.Cmd {
	return func() tea.Msg {
		_ = mgr.Stop(b.ID())
		return nil
	}
}

type TickMsg struct{}

func TickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*200, func(time.Time) tea.Msg {
		return TickMsg{}
	})
}

func ListenBotEvents(mgr *bot.Manager) tea.Cmd {
	return func() tea.Msg {
		ev := <-mgr.Events()
		return BotStateChangeMsg{
			ID:     ev.ID,
			Status: ev.Status,
		}
	}
}
