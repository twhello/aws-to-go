package services

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

var configSetting = &configuration{sync.RWMutex{}, 5, false}

type configuration struct {
	m             sync.RWMutex
	retryAttempts uint
	isDebug       bool
}

// Set the behavior of the shared http.Client.
// (maxIdleConnsPerHost uint) Controls the maximum idle (keep-alive) to keep per-host.
// (responseHeaderTimeout time.Duration) Specifies the amount of time to wait for a server's response headers after fully writing the request (including its body, if any).
func (c *configuration) SetHttpClient(maxIdleConnsPerHost uint, responseHeaderTimeout time.Duration) {
	c.m.Lock()
	httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   int(maxIdleConnsPerHost),
			ResponseHeaderTimeout: responseHeaderTimeout,
		},
	}
	c.m.Unlock()
}

// The number of times to retry a service request when returning a service or throttle exception.
func (c *configuration) RetryAttempts() uint {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.retryAttempts
}

// Set the number of retry attempts when a service request returns a service or throttle exceptio.
func (c *configuration) SetRetryAttempts(numRetries uint) {
	c.m.Lock()
	c.retryAttempts = numRetries
	c.m.Unlock()
}

// Returns the debugging flag setting.
func (c *configuration) IsDebugging() bool {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.isDebug
}

// Set the debugging flag to output log information.
func (c *configuration) SetDebugging(debug bool) {
	c.m.Lock()
	c.isDebug = debug
	c.m.Unlock()
}

// Returns the configuration settings for the service requests.
func Config() *configuration {
	return configSetting
}
