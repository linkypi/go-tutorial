/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2023-12-17 10:26:38
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2023-12-19 09:33:11
 * @FilePath: /test/interface.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main_interface() {

	var run Runner
	println(run)

	var u User
	u.setAge(25)
	(*User).setAge(&u, 25)

	Write(&File{})
	Write(&Memory{})

}

type Runner interface{}

type User struct {
	age int
}

func (u *User) setAge(age int) {
	u.age = age
}

func Write(w Writer) {
	fmt.Printf("%+v\n", w)
}

type Writer interface {
	write()
}

type File struct {
}

func (f *File) write() {

}

type Memory struct{}

func (f *Memory) write() {

}
