/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2024-01-04 12:26:43
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2024-01-04 12:30:46
 * @FilePath: /test/test/channel_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package test

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {

	defer func() {
		recover()
		fmt.Println("xxx...")
	}()
	var ch chan int
	ch <- 1
	// val := <-ch
	fmt.Println("test...")

}
