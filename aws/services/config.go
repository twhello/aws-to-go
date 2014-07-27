package services

import (
	"net/http"
	"sync"
	"time"
)

var configStatic = configuration{sync.RWMutex{}, 5}

type configuration struct {
	m             sync.RWMutex
	retryAttempts uint
}

func (c configuration) SetHttpClient(conns, timeout time.Duration) {
	c.m.Lock()
	httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   int(conns),
			ResponseHeaderTimeout: timeout,
		},
	}
	c.m.Unlock()
}

func (c configuration) RetryAttempts() uint {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.retryAttempts
}

func (c configuration) SetRetryAttempts(retries uint) {
	c.m.Lock()
	c.retryAttempts = retries
	c.m.Unlock()
}

func Config() *configuration {
	return &configStatic
}
