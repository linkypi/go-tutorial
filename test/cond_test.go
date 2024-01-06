package test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var cond sync.Cond

func produce(ch chan<- int, index int) {
	for {
		cond.L.Lock()
		// wait 会释放 cond.L 的锁, 并挂起当前协程
		// 直到 cond.Broadcast() 或 signal 被调用
		// 待wait 恢复后, cond.L 锁会被重新获得, 但是条件
		// 场景分析: 当前 ch 缓冲区已满, 无法再写入,应该等待消费者消费后才能写入
		// 此时有五个协程在争抢 cond.L 锁, 假如只有一个协程2可以获得锁, 待协程 2 运行到
		// 该位置时, 同样执行 cond.wait()释放 cond.L 锁, 然后挂起. 此时其他协程就可以在
		// cond.wait() 处被唤醒, 继续执行后面代码逻辑, 若此处不使用 for 判断, 那么其他协程
		// 就会在没有判断 ch 缓冲区是否已满的情况下继续往下执行. 但当前 ch 缓冲区已满, 无法写入, 只能死等
		for len(ch) == 3 {
			cond.Wait()
		}
		num := rand.Intn(100)
		ch <- num
		fmt.Printf("produce: %d\n", num)

		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Second)
	}
}

func consume(ch <-chan int, index int) {
	for {
		cond.L.Lock()
		for len(ch) == 0 {
			cond.Wait()
		}
		num := <-ch
		fmt.Printf("consume: %d\n", num)

		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Second)
	}
}

func TestCondition(t *testing.T) {
	//quit := make(chan struct{})
	produceCh := make(chan int, 3)
	cond.L = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go produce(produceCh, i+1)
	}
	for i := 0; i < 3; i++ {
		go consume(produceCh, i+1)
	}

	//<-quit
	select {}
}
