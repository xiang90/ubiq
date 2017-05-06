package work

import (
	"context"

	"github.com/xiang90/ubiq/ubiq"
)

type UnitCommitter struct {
}

func (uc *UnitCommitter) CommitWorkUnit(ctx context.Context, wu *ubiq.WorkUnit) error {
	return nil
}
