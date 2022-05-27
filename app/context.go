package app

import (
	"context"

	"github.com/NuttapolCha/simple-covid-grouping/client"
	"github.com/NuttapolCha/simple-covid-grouping/log"
)

type Context struct {
	ctx    context.Context
	logger log.Logger
	conn   client.Client
}

func NewAppContext(app *App) (*Context, error) {
	clientConfig, err := client.InitConfig()
	if err != nil {
		return nil, err
	}
	httpClient, err := client.NewHttpClient(clientConfig)
	if err != nil {
		return nil, err
	}

	return &Context{
		conn: httpClient,
	}, nil
}

func (ctx *Context) getContext() context.Context {
	return ctx.ctx
}
