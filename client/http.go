package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/NuttapolCha/simple-covid-grouping/log"

	"net/http"
)

type HttpClient struct {
	logger log.Logger

	c      *http.Client
	config *Config
}

func NewHttpClient(config *Config) (Client, error) {
	return &HttpClient{
		c: &http.Client{
			Timeout: config.requestTimeoutSec,
		},
		config: config,
	}, nil
}

func (client *HttpClient) Get(url string, placeHolder interface{}) (err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	attmps := 0
	for i := 0; i < client.config.retryCount+1; i++ {
		var resp *http.Response

		// do request
		resp, err = client.c.Do(req)
		if err != nil {
			client.logger.Errorf("attemp %d: failed to GET %s, because: %v, retrying in %v", attmps+1, url, err, client.config.timeBetweenRetry)
			time.Sleep(client.config.timeBetweenRetry)
			attmps++
			continue
		}
		err = determineHttpStatus(resp)
		if err != nil {
			client.logger.Errorf("attemp: %d could not GET %s because: %v, retrying in %v", attmps+1, url, err, client.config.timeBetweenRetry)
			time.Sleep(client.config.timeBetweenRetry)
			attmps++
			continue
		}

		// process response
		var body []byte
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			client.logger.Errorf("could not read response body because: %v", err)
			return err
		}
		if err := json.Unmarshal(body, placeHolder); err != nil {
			client.logger.Errorf("could not unmarshal JSON bytes to GO struct because: %v", err)
			return err
		}

		client.logger.Infof("requested to %s was successfully", url)
		return nil
	}

	return fmt.Errorf("could not GET %s after %d attempts, latest error: %v", url, attmps+1, err)
}

func determineHttpStatus(resp *http.Response) error {
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected response")
	}
	return nil
}
