package state

type logStateType int

var (
	logStateWaiting    = logStateType(1)
	logStateInProgress = logStateType(2)
	logStateCommitted  = logStateType(3)
)

type logState struct {
	fileName string
	beg      int64
	end      int64
	state    logStateType
}

func (a *logState) merge(b *logState) bool {
	if a.end != b.beg {
		return false
	}
	if a.state == b.state {
		a.end = b.end
		return true
	}
	return false
}
