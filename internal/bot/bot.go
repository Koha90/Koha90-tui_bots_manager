package bot

import "context"

type Status int

const (
	Stopped Status = iota
	Starting
	Running
	Error
)

var StatusNames = map[Status]string{
	Stopped:  "stopped",
	Starting: "starting",
	Running:  "running",
	Error:    "error",
}

type Bot interface {
	ID() string
	Start(ctx context.Context) error
	Stop() error
	Status() Status
}

func (s Status) String() string {
	if v, ok := StatusNames[s]; ok {
		return v
	}
	return "unknown"
}
