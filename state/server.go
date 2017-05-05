package state

import "errors"

var (
	ErrNotFound = errors.New("file not found")
)

type Server struct {
	logs map[string]logState
}

func (s *Server) Notify(e Event) error {
	switch e.Type {
	case EventNew:
		if _, ok := s.logs[e.FileName]; ok {
			return nil
		}
		s.logs[e.FileName] = logState{
			fileName: e.FileName,
			beg:      0,
			end:      e.Size,
		}
		// todo: persist to etcd
	case EventIncr:
		ls, ok := s.logs[e.FileName]
		if !ok {
			return ErrNotFound
		}
		if ls.end > e.Size {
			// duplicate notification
			return nil
		}
		nls := &logState{
			beg:   ls.end,
			end:   e.Size,
			state: logStateWaiting,
		}
		ls.merge(nls)
	}

	return nil
}
