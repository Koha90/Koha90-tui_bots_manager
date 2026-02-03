package bot

import "context"

type Status int

const (
	Stopped Status = iota
	Starting
	Running
	Error
)

type Bot interface {
	ID() string
	Start(ctx context.Context) error
	Stop() error
	Status() Status
}

func (s Status) String() string {
	switch s {
	case Running:
		return "running"
	case Starting:
		return "starting"
	case Stopped:
		return "stopped"
	case Error:
		return "error"
	default:
		return "unknown"
	}
}
