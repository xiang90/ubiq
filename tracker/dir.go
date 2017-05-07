package tracker

import (
	"io/ioutil"
	"path/filepath"
	"time"
)

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
	t          eventType
	origOffset int64
	offset     int64
	fileName   string
}

type dirTracker struct {
	dir   string
	files map[string]fileTracker
}

func (dt *dirTracker) track() (<-chan event, error) {
	// check if the directory exsits
	_, err := ioutil.ReadDir(dt.dir)
	if err != nil {
		return nil, err
	}

	eventC := make(chan event) // TODO size

	go func() {
		for {
			files, err := ioutil.ReadDir(dt.dir)
			if err != nil {
				// TODO log the error
				continue
			}
			for _, file := range files {
				fn := file.Name()
				// only track files with proper trackSuffix
				if filepath.Ext(fn) != trackSuffix {
					continue
				}
				if ft, ok := dt.files[fn]; ok {
					// generate a fileIncr event for exsiting files
					n := file.Size()
					if n > ft.offset {
						e := event{t: fileIncr, origOffset: ft.offset, offset: n, fileName: fn}
						ft.offset = n
						dt.files[fn] = ft
						eventC <- e
					}
				} else {
					// generate a fileNew event for new files
					n := file.Size()
					dt.files[fn] = fileTracker{offset: n}
					e := event{t: fileNew, origOffset: 0, offset: n, fileName: fn}
					eventC <- e
				}
			}
			// TODO make the size configurable
			time.Sleep(10 * time.Millisecond)
		}
	}()

	return eventC, nil
}
