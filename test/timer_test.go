package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(time.Second)
	now := <-timer.C
	fmt.Println("当前时间: ", now)

	time.AfterFunc(time.Second, func() {
		fmt.Println("after running...")
	})
}
