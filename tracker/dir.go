package tracker

const (
	trackSuffix = ".ubiq"
)

type fileTracker struct {
	offset int64
}

type eventType int

var (
	fileNew  = eventType(1)
	fileIncr = eventType(2)
)

type event struct {
	t      eventType
	offset int64
}

type dirTracker struct {
	dir   string
	files map[string]fileTracker
}

func (dt *dirTracker) track() (<-chan event, error) {

	return nil, nil
}
