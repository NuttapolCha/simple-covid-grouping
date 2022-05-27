package client

type Client interface {
	Get(url string, resp interface{}) error
}
