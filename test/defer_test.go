/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2024-01-04 11:45:52
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2024-01-04 12:26:30
 * @FilePath: /test/test/defer_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package test

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	// fmt.Println(f1())
	// fmt.Println(f2())
	fmt.Printf("f3: %d\n", *f3())
	fmt.Printf("f6: %d\n", f6())
	fmt.Println(f5())

	fmt.Println(test12())
	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3())
	fmt.Println(test4())

	return
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5 // 1.x = 5 2. defer x = 6 , 3  真正的返回
}

func f3() (x *int) {
	defer func() {
		*x++
	}()
	a := 5
	return &a // 返回前先赋值给 x, 然后执行 defer 直接修改指针地址数据而非副本数据
}
func f6() (x int) {
	defer func(x *int) *int {
		(*x)++
		return x
	}(&x)
	return 5 // 1. x = 5 // 2.defer  x =6  3. ret返回
}
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值=x
}

func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是x, 不是返回值, 所以返回结果不变
	}()
	return x // 1. 返回值赋值 2.defer 3.真正的ret指令
}

func test12() (result int) {
	defer func() {
		fmt.Println(result)
		result = result + 1
		// return 0
	}()
	// 返回前先将返回值赋值给返回变量 result
	return 2
}

func test1() (v int) {
	defer fmt.Println(v)
	return v
}

func test2() (v int) {
	defer func() {
		fmt.Println(v)
	}()
	return 3
}

func test3() (v int) {
	defer fmt.Println(v)
	v = 3
	return 4
}
func test4() (v int) {
	defer func(n int) {
		fmt.Println(n)
	}(v)
	return 5
}
