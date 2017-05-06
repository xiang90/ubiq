package work

import (
	"context"

	"github.com/xiang90/ubiq/ubiq"
)

type UnitRetriever struct {
}

func (uc *UnitRetriever) RetrieveWorkUnit(ctx context.Context) (*ubiq.WorkUnit, error) {
	return nil, nil
}
