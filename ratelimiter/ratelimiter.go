package ratelimiter

import (
	"sync"
	"time"
)

var mutex = &sync.Mutex{}
var cache = map[string]*RateLimiter{}

// Returns a new or cached RateLimiter with the specified stable throughput, given as "permits per second".
func CacheNew(key string, permitsPerSecond uint32) *RateLimiter {

	mutex.Lock()
	defer mutex.Unlock()

	limiter, ok := cache[key]
	if !ok {
		limiter = New(permitsPerSecond)
		limiter.key = key
		cache[key] = limiter
	}
	return limiter
}

// Creates a RateLimiter with the specified stable throughput, given as "permits per second".
func New(permitsPerSecond uint32) *RateLimiter {

	limiter := &RateLimiter{"NO_KEY", float64(permitsPerSecond), &sync.Mutex{}, 0, time.Now()} // make(chan aquireRequest), make(chan bool)}
	return limiter
}

// Remove the specified RateLimiter from the cache.
func CacheRemove(key string) {

	mutex.Lock()
	delete(cache, key)
	mutex.Unlock()
}

// Removes all RateLimiters from the cache.
func CacheRemoveAll() {

	mutex.Lock()
	defer mutex.Unlock()

	for k, _ := range cache {
		delete(cache, k)
	}
}

/******************************************************************************
Rate Limiter object.
*/
type RateLimiter struct {
	key  string
	pps  float64 // Permits Per Second
	m    *sync.Mutex
	cpps float64   // Consumed Permits Per Second
	last time.Time // Last aquire attempt
}

func (l *RateLimiter) Key() string {
	return l.key
}

func (l *RateLimiter) Rate() int {
	return int(l.pps)
}

// Acquires the given number of permits from this RateLimiter, blocking until the request can be granted.
func (l *RateLimiter) Aquire(permits int) {
	l.m.Lock()
	l.aquireRequest(uint32(permits), true)
	l.m.Unlock()
}

// Acquires a permit from this RateLimiter if it can be acquired immediately without delay.
func (l *RateLimiter) TryAquire(permits int) bool {
	l.m.Lock()
	defer l.m.Unlock()
	return l.aquireRequest(uint32(permits), false)
}

func (l *RateLimiter) aquireRequest(permits uint32, blocking bool) bool {

	now := time.Now()
	elapsed := now.Sub(l.last)
	l.last = now
	l.cpps -= l.pps * (float64(elapsed.Nanoseconds()) / 1e9)
	if l.cpps < 0 {
		l.cpps = 0
	}

	if !blocking && l.cpps+float64(permits) > l.pps {
		return false
	}

	l.cpps += float64(permits)
	if l.cpps < l.pps {
		return true
	}

	time.Sleep(time.Duration(((l.cpps - l.pps) / l.pps) * 1e9))

	return true
}
