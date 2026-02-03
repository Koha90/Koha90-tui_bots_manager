package bot

import (
	"context"
	"sync/atomic"
	"time"
)

type FakeBot struct {
	id     string
	status atomic.Int32

	cancel context.CancelFunc
}

func NewFake(id string) *FakeBot {
	b := &FakeBot{id: id}
	b.status.Store(int32(Stopped))
	return b
}

func (b *FakeBot) ID() string {
	return b.id
}

func (b *FakeBot) Status() Status {
	return Status(b.status.Load())
}

func (b *FakeBot) Start(ctx context.Context) error {
	// зашита от двойного старта
	if b.Status() == Running || b.Status() == Starting {
		return nil
	}

	ctx, cancel := context.WithCancel(ctx)
	b.cancel = cancel

	b.status.Store(int32(Starting))

	go func() {
		time.Sleep(500 * time.Millisecond)
		b.status.Store(int32(Running))

		<-ctx.Done()
		b.status.Store(int32(Stopped))
	}()

	return nil
}

func (b *FakeBot) Stop() error {
	if b.cancel != nil {
		b.cancel()
		b.cancel = nil
	}
	return nil
}
