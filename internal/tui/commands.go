package tui

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/koha90/tui_bots_manager/internal/bot"
)

func StartBotCmd(b bot.Bot) tea.Cmd {
	return func() tea.Msg {
		_ = b.Start(context.Background())
		return BotStateChangeMsg{}
	}
}

func StopBotCmd(b bot.Bot) tea.Cmd {
	return func() tea.Msg {
		_ = b.Stop()
		return BotStateChangeMsg{}
	}
}

type TickMsg struct{}

func TickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*200, func(time.Time) tea.Msg {
		return TickMsg{}
	})
}
