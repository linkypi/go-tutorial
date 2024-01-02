/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2023-12-27 14:31:39
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2024-01-01 09:42:05
 * @FilePath: /test/search.go
 * @Description: 要求使用 10 个 goroutine 来实现快速查询功能
 */
package tutorial

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const NUM_OF_WORKERS int = 10

func FindNumberWithWorkers(arr []int, target int) bool {
	return findNumberWithWorkersInternal(arr, target, 0)
}
func FindNumberWithWorkersAndTimeout(arr []int, target int, timeout time.Duration) bool {
	return findNumberWithWorkersInternal(arr, target, timeout)
}
func findNumberWithWorkersInternal(arr []int, target int, timeout time.Duration) bool {

	// count := len(mySlice)
	// 使用多个 goroutine 同时查找, 限制超时时间, 超出时间则自动执行 ctx.Done()
	var ctx context.Context
	var cancel context.CancelFunc
	if timeout == 0 {
		ctx, cancel = context.WithCancel(context.Background())
	} else {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	}
	defer cancel()

	var wGroup sync.WaitGroup
	wGroup.Add(NUM_OF_WORKERS)

	result := make(chan int)
	go FindWithWorkers(&wGroup, ctx, result, arr, target)

	var founded, finished bool
	var stopChan = make(chan struct{})

	go func(wGroup *sync.WaitGroup) {
		// 此处存在三种场景
		// 1. 元素已找到, 某个 worker 找到后会通过 result 管道通知主协程,
		//    然后执行 cancel 函数来取消其他协程的查询, 使得其他 worker 都可以顺利退出
		// 2. 元素不存在, 即所有协程都已查询完成但仍未发现目标, 随后所有 worker 会自动退出
		// 3. 查询超时, worker 中的 cancel.Done() 自动触发, 随后所有 worker 会自动退出
		wGroup.Wait()
		// wgroup 得知所有 worker 顺利退出后执行一下代码
		// fmt.Println("wait finished.", time.Now())
		// 此时需通知 main 协程查询已完成
		close(stopChan)
	}(&wGroup)

	for {
		select {
		case <-stopChan:
			// 查询已完成, 此处对应上方三种场景
			finished = true
		// 查询到结果
		case res := <-result:
			fmt.Printf("%s worker %d find result.\n", time.Now(), res)
			founded = true
			// 任意一个工作协程找到目标后则取消其他协程的查找, 执行 cancel 后
			// worker 协程收到相关信号并在退出前执行 wgroup.Done(),
			// 待所有 worker 都顺利退出后由 stopChan来通知主协程已查询结束
			cancel()

		default:
			if finished {
				// goto Loop
				break
			}
		}

		if finished {
			break
		}
	}

	// Loop:
	// time.Sleep(time.Second * 2)
	if founded {
		fmt.Print("target number founded.\n")
	} else {
		fmt.Print("target number not found.\n")
	}
	return founded
	// fmt.Println("main func release.")
}

func FindWithWorkers(wgroup *sync.WaitGroup, myCtx context.Context, result chan int, arr []int, target int) {
	total := len(arr)
	if total == 0 {
		result <- -1
		return
	}
	// 将数据均分到 NUM_OF_WORKERS 个 工作协程进行查询
	countPerRoutine := total / NUM_OF_WORKERS
	index := 0
	var replica []int
	for i := 0; i < NUM_OF_WORKERS; i++ {
		if i == 9 {
			replica = arr[index:]
		} else {
			replica = arr[index:(index + countPerRoutine)]
		}
		index += countPerRoutine
		// fmt.Printf("replica slice len: %d , goroutine: %d\n", len(replica), i)
		go FindEle(wgroup, myCtx, result, i, replica, target)
	}

}

func FindEle(wgroup *sync.WaitGroup, ctx context.Context,
	result chan int, index int, replica []int, target int) {

	// fmt.Printf("worker: %d starting, replica slice len: %d .\n", index, len(replica))

	defer func() {
		wgroup.Done()
		// fmt.Printf("worker %d wait done.\n", index)
	}()

	for _, v := range replica {
		// time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d search timeout.\n", index)
			return
		default:
			// fmt.Printf("worker %d searching...%d\n", index, v)
			if v == target {
				result <- index
				fmt.Printf("worker %d found it. \n", index)
				return
			}
		}
	}
	// fmt.Printf("%s worker %d exit. \n", time.Now(), index)
}
