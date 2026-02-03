package tui

type StartBotMsg struct {
	ID string
}

type StopBotMsg struct {
	ID string
}

type (
	LoadBotsMsg   struct{}
	BotsLoadedMsg struct {
		Bots []BotView
	}
)
