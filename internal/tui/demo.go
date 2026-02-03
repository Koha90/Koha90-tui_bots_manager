package tui

func DemoBots() []BotView {
	return []BotView{
		{ID: "shop_1", Status: Running},
		{ID: "shop_2", Status: Stopped},
		{ID: "shop_3", Status: Error},
	}
}
