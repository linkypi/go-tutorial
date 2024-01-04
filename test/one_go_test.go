/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2024-01-04 15:35:02
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2024-01-04 15:38:22
 * @FilePath: /test/test/one_go_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestOneGoroutine(t *testing.T) {
	// 模拟单核 CPU
	runtime.GOMAXPROCS(1)

	// 模拟 Goroutine 死循环
	go func() {
		for {
		}
	}()
	go func() {
		for {
		}
	}()
	time.Sleep(time.Second)
	fmt.Println("脑子进煎鱼了")
}
