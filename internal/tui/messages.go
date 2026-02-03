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

type BotStateChangeMsg struct{}
