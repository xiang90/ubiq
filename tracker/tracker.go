package tracker

type DirServer struct {
	dirs []dirTracker
}

func NewDirServer(root string) (*DirServer, error) {
	s := &DirServer{dirs: make([]dirTracker, 0)}
	return s, nil
}

func (s *DirServer) Start() error {
	return nil
}

func (s *DirServer) AddDir(dir string) (<-chan event, error) {
	dt := dirTracker{dir: dir}
	ch, err := dt.track()
	if err != nil {
		return nil, err
	}
	s.dirs = append(s.dirs, dt)
	return ch, nil
}

type S3Server struct {
}
