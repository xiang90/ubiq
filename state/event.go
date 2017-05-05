package state

type EventType int

var (
	EventNew  = EventType(1)
	EventIncr = EventType(2)
)

type Event struct {
	Type     EventType
	FileName string
	Size     int64
}
