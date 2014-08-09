package ratelimiter

import (
	"fmt"
	"runtime"
    "testing"
    "time"
)

func TestRateLimiter(t *testing.T) {

	runLimiter := func(num int) {
		
		rl := CacheNew("test", 1000)
		
		for i := 0 ; i < 10 ; i++ {
			fmt.Println("TryAquire", num, rl.TryAquire(100))
			rl.Aquire(100*num)
			fmt.Println("Acquire", num, i)
			runtime.Gosched()
		}
		fmt.Println(">>> End", num)
	}
	
	go runLimiter(1)
	go runLimiter(2)
	go runLimiter(3)
	go runLimiter(4)
	
	time.Sleep(20 * 1e9)

	fmt.Println(">>> End Test")
}

