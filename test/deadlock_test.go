package test

import (
	"fmt"
	"sync"
	"testing"
)

func TestDeadLock1(t *testing.T) {
	var wg sync.WaitGroup
	foo := make(chan int)
	bar := make(chan int)
	closing := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		// foo <- <-bar 首先从bar获取数据,然后放入到 foo,
		// 但是由于 bar 没有输入, 一直被堵塞, 最后出现死锁
		// 若在主协程将数据写入 bar 则死锁消除: bar <- 123
		case foo <- <-bar:
		case <-closing:
			println("closing")
		}
	}()
	//bar <- 123
	close(closing)
	wg.Wait()
}

func TestDeadLock2(t *testing.T) {

	ch := make(chan int)
	// 无缓冲 channel 写入后没有消费导致阻塞, 消费端
	ch <- 1

	a := <-ch
	fmt.Println("num:", a)
}
