package tracker

type DirServer struct {
	dirs []dirTracker
}

func NewDirServer(root string) (*DirServer, error) {
	return nil, nil
}

func (s *DirServer) Start() error {
	return nil
}

type S3Server struct {
}
