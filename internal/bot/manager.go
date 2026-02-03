// Package bot - Start / Stop / Status
package bot

import (
	"context"
	"errors"
	"sync"
)

type ManagedBot struct {
	Bot
	cancel context.CancelFunc
}

type Manager struct {
	mu   sync.Mutex
	bots map[string]*ManagedBot
}

func NewManager() *Manager {
	return &Manager{
		bots: make(map[string]*ManagedBot),
	}
}

func (m *Manager) Add(b Bot) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.bots[b.ID()] = &ManagedBot{Bot: b}
}

func (m *Manager) Start(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	b, ok := m.bots[id]
	if !ok {
		return errors.New("bot not found")
	}

	if b.cancel != nil {
		return nil // уже запущен
	}

	ctx, cancel := context.WithCancel(context.Background())
	b.cancel = cancel

	return b.Start(ctx) // TODO: b.Bot.Start
}

func (m *Manager) Stop(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	b, ok := m.bots[id]
	if !ok || b.cancel == nil {
		return nil
	}

	b.cancel()
	b.cancel = nil
	return b.Stop() // TODO: b.BOT.Stop()
}

func (m *Manager) List() []Bot {
	m.mu.Lock()
	defer m.mu.Unlock()

	out := make([]Bot, 0, len(m.bots))
	for _, b := range m.bots {
		out = append(out, b.Bot)
	}

	return out
}
