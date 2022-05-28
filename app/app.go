package app

import (
	"github.com/NuttapolCha/simple-covid-grouping/client"
	"github.com/NuttapolCha/simple-covid-grouping/log"
)

type App struct {
	Logger log.Logger
	conn   client.Client
}

func New(logger log.Logger) (*App, error) {
	clientConfig, err := client.InitConfig()
	if err != nil {
		return nil, err
	}
	httpClient, err := client.NewHttpClient(clientConfig, logger)
	if err != nil {
		return nil, err
	}

	application := &App{
		Logger: logger,
		conn:   httpClient,
	}

	return application, nil
}
