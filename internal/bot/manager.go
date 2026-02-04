// Package bot - Start / Stop / Status
package bot

import (
	"context"
	"sync"
	"time"
)

type ManagedBot struct {
	Bot
	cancel context.CancelFunc
}

type Manager struct {
	mu     sync.Mutex
	bots   map[string]*ManagedBot
	events chan BotEvent
}

type BotEvent struct {
	ID     string
	Status Status
}

func NewManager() *Manager {
	return &Manager{
		bots:   make(map[string]*ManagedBot),
		events: make(chan BotEvent, 16),
	}
}

func (m *Manager) Add(b Bot) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.bots[b.ID()] = &ManagedBot{Bot: b}
}

func (m *Manager) Start(id string) error {
	m.mu.Lock()
	b, ok := m.bots[id]
	if !ok || b.cancel != nil {
		m.mu.Unlock()
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	b.cancel = cancel
	m.mu.Unlock()

	_ = b.Start(ctx)

	go func() {
		for {
			status := b.Status()
			m.events <- BotEvent{ID: id, Status: status}
			if status == Stopped || status == Error {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	return nil
}

func (m *Manager) Stop(id string) error {
	m.mu.Lock()
	b, ok := m.bots[id]
	if !ok || b.cancel == nil {
		m.mu.Unlock()
		return nil
	}

	b.cancel()
	b.cancel = nil
	m.mu.Unlock()

	_ = b.Stop()
	m.events <- BotEvent{ID: id, Status: b.Status()}
	return nil
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

func (m *Manager) Events() <-chan BotEvent {
	return m.events
}

func (m *Manager) Status(id string) Status {
	m.mu.Lock()
	defer m.mu.Unlock()

	b, ok := m.bots[id]
	if !ok {
		return Stopped
	}
	return b.Status()
}
