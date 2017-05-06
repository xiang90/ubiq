package state

import "errors"

var (
	ErrNotFound = errors.New("file not found")
)

type Server struct {
	logs map[string]*logStates
}

// Notify notifies the state server with an event from the log data tracker.
func (s *Server) Notify(e Event) error {
	switch e.Type {
	case EventNew:
		if _, ok := s.logs[e.FileName]; ok {
			return nil
		}
		s.logs[e.FileName] = newLogStates(e.FileName, e.Size)
		// todo: persist to etcd
	case EventIncr:
		ls, ok := s.logs[e.FileName]
		if !ok {
			return ErrNotFound
		}
		if ok := ls.merge(e); ok {
			// todo: persist to etcd
		}
	}

	return nil
}

func (s *Server) CreateWorkUnit() {

}

func (s *Server) RetrieveWorkUnit() {

}

func (s *Server) CommitWorkUnit() {

}
