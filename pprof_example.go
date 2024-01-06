/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2024-01-04 17:16:29
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2024-01-04 18:40:20
 * @FilePath: /test/test/timer_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// 在使用图形化方式查看性能时需先安装 Graphviz: http://www.graphviz.org/download/

// 性能分析流程:
// 1. import 导入 _ "net/http/pprof", 并在代码启动监听 localhost:6060
// 2. 启动完成后在页面查看 pprof 页面: http://localhost:6060/debug/pprof/

//  1. 使用 web 图表查看内存分配, 首先保存相关文件, 然后启动 web 页面查看
// go tool pprof http://localhost:6060/debug/pprof/allos 进入pprof命令行后可以直接使用 web 通过浏览器查看
// 若需要在浏览器查看所有指标信息则需要在命令行执行, 其中 xxx.inuse_space.001.pb.gz 由上方命令生成:
// go tool pprof -http=localhost:8081 /Users/leo/pprof/pprof.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz

// 2. 使用 web 图表查看 goroutine , 首先保存相关文件, 然后启动 web 页面查看
// go tool pprof http://localhost:6060/debug/pprof/goroutine 进入pprof命令行后可以直接使用 web 通过浏览器查看
// go tool pprof -http=localhost:8082 /Users/leo/pprof/pprof.goroutine.001.pb.gz
func main() {

	go func() {
		http.ListenAndServe("localhost:6060", nil)
		fmt.Println("server started.")
	}()

	timer := time.NewTimer(3 * time.Minute)
	defer timer.Stop()

	ch := make(chan int, 10)
	go func() {
		in := 1
		for {
			in++
			ch <- in
		}
	}()

	for {
		select {
		case _ = <-ch:
			// do something...
			continue
		// 在 case 中使用 time.After 会导致 timer 不断被创建, 进而出现 CPU 飙升, 内存飙升问题
		// case <-time.After(3 * time.Minute):
		// 	fmt.Printf("现在是：%d,我脑子进煎鱼了！", time.Now().Unix())
		case <-timer.C:
			fmt.Printf("现在是：%d，我脑子进煎鱼了！", time.Now().Unix())

		}
	}

}
