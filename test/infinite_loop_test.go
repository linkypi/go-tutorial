package test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestInfiniteLoop(t *testing.T) {
	var x int
	// 所有资源都让死循环占用,导致主协程无法执行
	numOfGoroutines := runtime.GOMAXPROCS(1)
	for i := 0; i < numOfGoroutines; i++ {
		go func() {
			for {
				x++
			}
		}()
	}

	time.Sleep(time.Millisecond)
	fmt.Printf("x = %v", x)
}
