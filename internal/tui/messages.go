package tui

import "github.com/koha90/tui_bots_manager/internal/bot"

type StartBotMsg struct {
	ID string
}

type StopBotMsg struct {
	ID string
}

type (
	LoadBotsMsg   struct{}
	BotsLoadedMsg struct {
		Bots []bot.Bot
	}
)

const (
	ConfirmStopMsg = "⚠ press x again to stop"
	ErrorBotMsg    = "❌ bot error, check logs"
	AlreadyRunning = "▶️ bot already running"
	AlreadyStopped = "⏹️bot already stopped"
)
