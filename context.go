/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2023-09-16 13:07:58
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2023-09-16 14:30:28
 * @FilePath: /test/context.go
 */

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main5() {

	req, _ := http.NewRequest("GET", "https://api.github.com/users/helei112g", nil)
	// 这里设置了超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("request Err", err.Error())
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	log.Default().Println("get info from url success.")
	go Hello()
	fmt.Println("jello")
}

func Hello() {
	time.Sleep(5 * time.Second)
	fmt.Println("hello everybody , I'm lineshen")
}

type key int

const (
	userIP = iota
	userID
	logID
)

type Result struct {
	order     string
	logistics string
	recommend string
}

// timeout: 1s
// 入口函数
func api() (result *Result, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	// 设置值
	ctx = context.WithValue(ctx, userIP, "127.0.0.1")
	ctx = context.WithValue(ctx, userID, 666888)
	ctx = context.WithValue(ctx, logID, "123456")
	result = &Result{}
	// 业务逻辑处理放到协程中
	go func() {
		result.order, err = getOrderDetail(ctx)
	}()
	go func() {
		result.logistics, err = getLogisticsDetail(ctx)
	}()
	go func() {
		result.recommend, err = getRecommend(ctx)
	}()
	for {
		select {
		case <-ctx.Done():
			return result, ctx.Err() // 取消或者超时，把现有已经拿到的结果返回
		default:
		}
		// 有错误直接返回
		if err != nil {
			return result, err
		}
		// 全部处理完成，直接返回
		if result.order != "" && result.logistics != "" && result.recommend != "" {
			return result, nil
		}
	}
}

// timeout: 500ms
func getOrderDetail(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()
	// 模拟超时
	time.Sleep(time.Millisecond * 700)
	// 获取 user id
	uip := ctx.Value(userIP).(string)
	fmt.Println("userIP", uip)
	return handleTimeout(ctx, func() string {
		return "order"
	})
}

// timeout: 700ms
func getLogisticsDetail(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*700)
	defer cancel()
	// 获取 user id
	uid := ctx.Value(userID).(int)
	fmt.Println("userID", uid)
	return handleTimeout(ctx, func() string {
		return "logistics"
	})
}

// timeout: 400ms
func getRecommend(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*400)
	defer cancel()
	// 获取 log id
	lid := ctx.Value(logID).(string)
	fmt.Println("logID", lid)
	return handleTimeout(ctx, func() string {
		return "recommend"
	})
}

// 超时的统一处理代码
func handleTimeout(ctx context.Context, f func() string) (string, error) {
	// 请求之前先去检查下是否超时
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	str := make(chan string)
	go func() {
		// 业务逻辑
		str <- f()
	}()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case ret := <-str:
		return ret, nil
	}
}
