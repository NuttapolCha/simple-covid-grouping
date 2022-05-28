package client

import "context"

type Client interface {
	Get(ctx context.Context, url string, resp interface{}) error
}
