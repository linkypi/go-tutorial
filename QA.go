/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2023-12-19 09:32:49
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2023-12-27 14:31:50
 * @FilePath: /test/QA.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

// type T struct {
// 	n int
// }

// const (
// 	x = iota
// 	_
// 	y
// 	z = "zz"
// 	k
// 	p = iota
// )

// func hello(i *int) {
// 	fmt.Println(*i)
// }

// type S struct {
// }

// func f(x interface{}) {
// }

// func g(x *interface{}) {
// }

// type MyInt int

// func (i MyInt) PrintInt() {
// 	fmt.Println(i)
// }

// type Math struct {
// 	x, y int
// }

// var m = map[string]*Math{
// 	"foo": &Math{2, 3},
// }

// var c = make(chan int)
// var a int

// func fx() {
// 	a = 1
// 	fmt.Println("fx running...")
// 	<-c
// 	fmt.Println("fx return...")
// }

// type UserAges struct {
// 	ages map[string]int
// 	sync.Mutex
// }

// func (ua *UserAges) Add(name string, age int) {
// 	ua.Lock()
// 	defer ua.Unlock()
// 	ua.ages[name] = age
// }

// func (ua *UserAges) Get(name string) int {
// 	if age, ok := ua.ages[name]; ok {
// 		return age
// 	}
// 	return -1
// }

// func A() int {
// 	time.Sleep(100 * time.Millisecond)
// 	return 1
// }

// func B() int {
// 	time.Sleep(1000 * time.Millisecond)
// 	return 2
// }

// type T1 int

// func F(t T1) {}

// type T2 []int

// func F2(t T2) {}

// const (
// 	azero = iota
// 	aone  = iota
// )

// const (
// 	info  = "msg"
// 	bzero = iota
// 	bone  = iota
// )

// func alwaysFalse() bool {
// 	return false
// }

// func fp() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Printf("recover:%#v", r)
// 		}
// 	}()
// 	panic(1)
// 	panic(2)
// }

// var fxb = func(i int) {
// 	print("x")
// }

// type Orange struct {
// 	Quantity int
// }

// func (o *Orange) Increase(n int) {
// 	o.Quantity += n
// }

// func (o *Orange) Decrease(n int) {
// 	o.Quantity -= n
// }

// func (o Orange) String() string {
// 	return fmt.Sprintf("%#v", o.Quantity)
// }
// func test(x int) (func(), func()) {
// 	return func() {
// 			println(x)
// 			x += 10
// 		}, func() {
// 			println(x)
// 		}
// }

// func F5(n int) func() int {
// 	return func() int {
// 		n++
// 		return n
// 	}
// }

type T struct{}

func (*T) foo() {
}

func (T) bar() {
}

type S struct {
	*T
}

func main_qa() {
	s := S{}
	fmt.Printf("%#v", s)
	_ = s.foo
	s.foo()
	_ = s.bar
}

// var k = 9
// for k = range []int{} {nbkjbhgciugfcxx
// 	fmt.Println("nnn")
// }
// fmt.Println(k)

// for k = 0; k < 3; k++ {
// }
// fmt.Println(k)

// for k = range (*[3]int)(nil) {
// 	fmt.Println(k)
// }
// fmt.Println(k)

// var x int8 = -128
// var y = x / -1
// fmt.Println(y)

// f := F5(5)
// defer func() {
// 	fmt.Println(f())
// }()
// defer fmt.Println(f())
// i := f()
// fmt.Println(i)

// defer func() {
// 	fmt.Println("xx")
// 	fmt.Print(recover())
// }()
// defer func() {
// 	defer func() {
// 		fmt.Println("nested..")
// 		fmt.Print(recover())
// 	}()
// 	fmt.Println("out..")
// 	panic(1)
// }()
// defer recover()
// panic(2)

// ts := [2]T{}
// for i, t := range &ts {
// 	switch i {
// 	case 0:
// 		t.n = 3
// 		ts[1].n = 9
// 	case 1:
// 		fmt.Print(t.n, " ")
// 	}
// }
// fmt.Println(ts)

// ts2 := [2]T{}
// for i := range ts2[:] {
// 	switch t := &ts2[i]; i {
// 	case 0:
// 		t.n = 3
// 		ts2[1].n = 9
// 	case 1:
// 		fmt.Print(t.n, " ")
// 	}
// }
// fmt.Println(ts2)

// a, b := test(100)
// a()
// b()

// var orange Orange
// orange.Increase(10)
// orange.Decrease(5)
// fmt.Println(orange)

// fxb := func(i int) {
// 	println(i)
// 	if i > 0 {
// 		fxb(i - 1)
// 	}
// }
// fxb(10)

// fp()

// switch alwaysFalse(); {
// case true:
// 	println(true)
// case false:
// 	println(false)
// }

// // fmt.Println(azero, aone)
// // fmt.Println(bzero, bone)

// // var q int
// // F(q)
// var q2 []int
// F2(q2)

// ch := make(chan int, 1)
// go func() {
// 	select {
// 	case ch <- A():
// 	case ch <- B():
// 	default:
// 		ch <- 3
// 	}
// }()
// fmt.Println(<-ch)

// func mainTest() {
// 	count := 1000
// 	gw := sync.WaitGroup{}
// 	gw.Add(count * 3)
// 	u := UserAges{ages: map[string]int{}}
// 	add := func(i int) {
// 		u.Add(fmt.Sprintf("user_%d", i), i)
// 		gw.Done()
// 	}
// 	for i := 0; i < count; i++ {
// 		go add(i)
// 		go add(i)
// 	}
// 	for i := 0; i < count; i++ {
// 		go func(i int) {
// 			defer gw.Done()
// 			u.Get(fmt.Sprintf("user_%d", i))
// 		}(i)
// 	}
// 	gw.Wait()
// 	fmt.Println("Done")
// }

// func mainxxx() {
// 	go fx()

// 	fmt.Println("main running...")
// 	c <- 0
// 	print(a)

// 	// arr := make([]int, 3, 4)
// 	// arr[0] = 0
// 	// arr[1] = 1
// 	// arr[2] = 2
// 	// fmt.Printf("main before: len: %d cap:%d data:%+v\n", len(arr), cap(arr), arr)
// 	// ap1(arr)
// 	// fmt.Printf("main ap1 after: len: %d cap:%d data:%+v\n\n", len(arr), cap(arr), arr)
// }
// func ap1(arr []int) {
// 	fmt.Printf("ap1 before:  len: %d cap:%d data:%+v\n", len(arr), cap(arr), arr)
// 	arr[0] = 11
// 	arr = append(arr, 111)
// 	fmt.Printf("ap1 after:  len: %d cap:%d data:%+v\n", len(arr), cap(arr), arr)
// }

// func main244() {

// 	v := []int{1, 2, 3}
// 	for i := range v {
// 		v = append(v, i)
// 	}

// 	// m["foo"].x = 4
// 	// fmt.Println(m["foo"].x)

// 	// s1 := []int{1, 2, 3}
// 	// s2 := s1[1:]
// 	// s2[1] = 4
// 	// fmt.Println(s1)
// 	// fmt.Println(s2)
// 	// s2 = append(s2, 5, 6, 7)
// 	// fmt.Println(s1)
// 	// fmt.Println(s2)

// 	s := S{}
// 	p := &s
// 	f(s) //A
// 	// g(s) //B
// 	f(p) //C
// 	// g(p) //D

// 	// var s1 []int
// 	// var s2 = []int{}
// 	// if s2 == nil {
// 	// 	fmt.Println("s2 is nil")
// 	// } else if s1 == nil {
// 	// 	fmt.Println("s1 is nil")
// 	// } else {
// 	// 	fmt.Println("no nil")
// 	// }

// 	i := 5
// 	defer hello(&i)
// 	i = i + 10
// 	// string 无法赋值为 nil
// var x string = nil

// a := [5]int{1, 2, 3, 4, 5}
// t := a[3:4:4]
// fmt.Println(t[0])

// fmt.Println(x, y, z, k, p)

// s := make([]int, 5)
// s = append(s, 1, 2, 3)
// fmt.Println(s)

// m := make(map[int]T)
// m[0] = T{2}
// fmt.Println(m[0].n)
// }
