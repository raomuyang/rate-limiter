package main

import (
	"github.com/rate-limiter/ratelimiter"
	"fmt"
)

func main() {
	tokenBucket()
}

func tokenBucket()  {
	limiter, _ := ratelimiter.CreateTokenBucket(1)
	for {
		elapsed, _ := limiter.Acquire(10)
		fmt.Printf("Wait nanos: %d\n", elapsed)
	}
}
