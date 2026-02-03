// Package bot - Start / Stop / Status
package bot

type Status int

const (
	Stopped Status = iota
	Running
	Error
)

type BotInfo struct {
	ID     string
	Status Status
}

type Manager struct {
	bots []BotInfo
}

func NewManager() *Manager {
	return &Manager{
		bots: []BotInfo{
			{ID: "shop_berlin", Status: Running},
			{ID: "shop_global", Status: Stopped},
			{ID: "test_bot", Status: Error},
		},
	}
}

func (m *Manager) List() []BotInfo {
	return m.bots
}
