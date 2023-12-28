/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2023-12-16 11:30:50
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2023-12-17 10:49:53
 * @FilePath: /test/value-reference.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"testing"
)

type Person struct {
	age int
}

func (p Person) howOld() int {
	return p.age
}

func (p *Person) GetAge2() {
	p.age += 1
}

func (p Person) GetAge() int {
	return p.age
}

func TestSyncmap(t *testing.T) {

}

func main_val() {
	// p1 是值类型
	p := Person{age: 18}

	// 值类型 调用接收者也是值类型的方法
	fmt.Println(p.howOld())

	// 值类型 调用接收者是指针类型的方法
	p.GetAge2()
	fmt.Println(p.GetAge())

	// ----------------------

	// p2 是指针类型
	p2 := &Person{age: 100}

	// 指针类型 调用接收者是值类型的方法
	fmt.Println(p2.GetAge())

	// 指针类型 调用接收者也是指针类型的方法
	p2.GetAge2()
	fmt.Println(p2.GetAge())
}
