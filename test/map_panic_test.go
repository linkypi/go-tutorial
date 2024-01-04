/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2023-12-28 16:46:46
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2024-01-04 11:16:41
 * @FilePath: /test/MyConcurrentMap.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package test

import (
	"fmt"
	"testing"
)

func TestMapConcurrentWritePanic(t *testing.T) {

	m := map[int]bool{}
	go func() {
		defer func() {
			a := recover()
			fmt.Println("1", a)
		}()
		for {
			m[10] = true
		}
	}()
	go func() {
		defer func() {
			a := recover()
			fmt.Println("2", a)
		}()
		for {
			m[10] = true
		}
	}()
	select {}

}
