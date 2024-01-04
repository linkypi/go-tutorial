/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2024-01-04 16:54:44
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2024-01-04 17:11:52
 * @FilePath: /test/test/escape_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package test

import (
	"fmt"
	"testing"
)

func TestEscape(t *testing.T) {
	a := new(struct{})
	b := new(struct{})
	println(a, b, a == b)

	c := new(struct{})
	d := new(struct{})
	// fmt.Println 引发变量逃逸到堆上, 即所有逃逸到堆上的变量都指向了 runtime.zerobase
	// 逃逸分析 go build -gcflags '-m -l' escape_test.go
	fmt.Println(c, d)
	println(c, d, c == d)
}
