package ubiq

import "context"

type Client struct {
}

func (c *Client) Consume(ctx context.Context) (*WorkUnit, error) {
	return nil, nil
}

func (c *Client) Commit(ctx context.Context, wb *WorkUnit, output []byte) error {
	return nil
}
