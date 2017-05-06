package state

type logStateType int

var (
	logStateWaiting    = logStateType(1)
	logStateInProgress = logStateType(2)
	logStateCommitted  = logStateType(3)
)

type logState struct {
	beg   int64
	end   int64
	state logStateType
}

type logStates struct {
	fileName string

	states []*logState
}

func newLogStates(name string, end int64) *logStates {
	return &logStates{
		fileName: name,
		states:   []*logState{{end: end, state: logStateWaiting}},
	}
}

func (ls *logStates) merge(e Event) bool {
	last := ls.last()

	if last.end > e.Size {
		// duplicate notification
		return false
	}

	if last.state == logStateWaiting {
		last.end = e.Size
	} else {
		ls.states = append(ls.states, &logState{beg: last.end + 1, end: e.Size, state: logStateWaiting})
	}
	return true
}

func (ls *logStates) last() *logState {
	return ls.states[len(ls.states)-1]
}
