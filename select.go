/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2023-12-16 09:47:24
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2023-12-16 11:31:01
 * @FilePath: /test/select.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main_select() {

	work := make(chan int, 100)

	stopChannel := make(chan struct{})

	waitGroup.Add(3)
	worker(work, stopChannel)
	worker(work, stopChannel)
	worker(work, stopChannel)

	for i := 0; i < 20; i++ {
		work <- i
	}

	time.Sleep(time.Second * 8)
	close(stopChannel)

	waitGroup.Wait()
	fmt.Print("all workers shutdown gracefully.\n")
	// time.Sleep(time.Second * 10)
}

func worker(work <-chan int, stopCh <-chan struct{}) {
	go func() {
		defer func() {
			fmt.Println("worker exit")
			waitGroup.Done()
		}()
		// Using stop channel explicit exit
		for {
			select {
			case <-stopCh:
				fmt.Println("Recv stop signal")
				return
			case x := <-work:
				fmt.Printf("Working %v.\n", x)
				time.Sleep(time.Second * 1)
			}
		}
	}()
}
