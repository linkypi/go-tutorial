package main

import (
	"context"
	"fmt"
	"time"
)

func createCtxWithVal1(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "name", "trouble")
	return child
}

func createCtxWithVal2(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "age", 20)
	return child
}

func printfValue(ctx context.Context) {
	fmt.Printf("name: %s\n", ctx.Value("name"))
	fmt.Printf("age: %d\n", ctx.Value("age"))
}

func TestWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go runTask(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	fmt.Println("task has canceled")

	time.Sleep(30 * time.Second)
}

func main6() {

	TestWithCancel()
	// testCtxWithTimeout()

	// context 继承关系,
	// grandpa := context.TODO()
	// father := createCtxWithVal1(grandpa)
	// grandson := createCtxWithVal2(father)
	// printfValue(grandson)

	// fmt.Println("hello")
}

func testCtxWithTimeout() {
	ctx, cancel := context.WithTimeout(context.TODO(), 1000*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf("channel has closed: %v", err)
	}

}

func runTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("task canceled")
			return
		default:
			fmt.Println("task is running...")
			time.Sleep(1 * time.Second)
		}
	}
}
