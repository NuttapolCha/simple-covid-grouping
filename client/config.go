package client

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	requestTimeoutSec time.Duration
	timeBetweenRetry  time.Duration
	retryCount        int
}

func InitConfig() (*Config, error) {
	return &Config{
		requestTimeoutSec: viper.GetDuration("Client.HTTP.RequestTimeoutSec") * time.Second,
		timeBetweenRetry:  viper.GetDuration("Client.HTTP.TimeBetweenRetryMilliSec") * time.Millisecond,
		retryCount:        viper.GetInt("Client.HTTP.RetryCount"),
	}, nil
}
