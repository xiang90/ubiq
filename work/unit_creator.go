package work

import "context"

// UnitCreator runs inside state server as a background go routine.
// Its goal is to convert the continuously growing log files into discrete work units.
// The Work Unit Creator maintains the maximum offset up to which the file has grown at
// each of the input logs data centers. It also stores the offset up to which work units
// have been created in the past for this file. As it creates new work units, it atomically
// updates the offset to ensure that each input byte is part of exactly one work unit.
// In order to prevent starvation, the Work Unit Creator prioritizes bytes from the oldest
// file while creating work units. The Work Unit Creator also tries to ensure that a work
// unit has chunks from several different files, as these could be read in parallel by
// the application.
type UnitCreator struct {
}

// Start starts unit creator as a background go routine.
func (uc *UnitCreator) Start(ctx context.Context) error {

}

func (uc *UnitCreator) createWorkUnit() {

}
