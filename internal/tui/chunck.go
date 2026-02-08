package tui

import "github.com/koha90/tui_bots_manager/internal/bot"

func chunkBots(bots []bot.Bot, size int) [][]bot.Bot {
	var chunks [][]bot.Bot

	for i := 0; i < len(bots); i += size {
		end := min(i+size, len(bots))
		chunks = append(chunks, bots[i:end])
	}

	return chunks
}
