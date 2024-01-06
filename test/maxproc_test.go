package test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMaxProcess(t *testing.T) {
	// 返回修改前的处理器数量
	gomaxprocs := runtime.GOMAXPROCS(8)
	fmt.Println(gomaxprocs)

	// 返回修改前的处理器数量 输出 8
	gomaxprocs = runtime.GOMAXPROCS(2)
	fmt.Println(gomaxprocs)

	// 返回修改前的处理器数量 输出 2
	gomaxprocs = runtime.GOMAXPROCS(1)
	fmt.Println(gomaxprocs)
}
