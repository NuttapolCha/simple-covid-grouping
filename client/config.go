package client

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	requestTimeout    time.Duration
	timeBetweenRetry  time.Duration
	retryCount        int
	allowUnsecureCall bool
}

func InitConfig() (*Config, error) {
	return &Config{
		requestTimeout:    viper.GetDuration("Client.HTTP.RequestTimeoutSec") * time.Second,
		timeBetweenRetry:  viper.GetDuration("Client.HTTP.TimeBetweenRetryMilliSec") * time.Millisecond,
		retryCount:        viper.GetInt("Client.HTTP.RetryCount"),
		allowUnsecureCall: viper.GetBool("Client.HTTP.AllowUnsecureCall"),
	}, nil
}
