/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2019-11-03 12:16:33
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2023-12-16 09:49:17
 * @FilePath: /test/a.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
)

type MyType int32

func main_defer() {

	result := f12()
	fmt.Printf("%+v", result)
	fmt.Printf("%+v", f1())
	fmt.Printf("%+v", f2())
	fmt.Printf("%+v", f3())
	fmt.Printf("%+v", f4())
	fmt.Printf("%+v", f5())
	fmt.Printf("%+v", f6())

	// newFunction()

	// time.Sleep(time.Second * 3)
}

func newFunction() {
	myType := new(MyType)

	name := make(chan int, 5)
	for i := 1; i < 5; i++ {
		name <- i
	}

	fmt.Printf("%+v\n\n", myType)
	fmt.Println("jello")

	go func(x chan int) {
		for {
			select {
			case a := <-x:
				fmt.Printf("%+v\n", a)
			default:
			}
		}
	}(name)
}

func f12() (result int) {
	defer func() {
		result = result + 1
		// return 0;
	}()
	return 2
}

// Go语言中函数的return不是原子操作，在底层是分为两步来执行
// 第一步：返回值赋值
// defer
// 第二步：真正的RET返回
// 函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间

func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是x不是返回值
	}()
	return x // 1. 返回值赋值 2.defer 3.真正的ret指令
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值=x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ // 修改的是x
	}()
	return x // 1. 返回值 = y = x = 5 2. defer修改的是x 3. 真正的返回
}

func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是函数中x的副本
	}(x)
	return 5 // 返回值 = x = 5
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5 // 1.x = 5 2. defer x = 6 3  真正的返回
}

func f6() (x int) {
	defer func(x *int) *int {
		(*x)++
		return x
	}(&x)
	return 5 // 1. x = 5 // 2.defer  x =6  3. ret返回
}
