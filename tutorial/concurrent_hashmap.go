/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2023-12-28 16:46:46
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2024-01-01 11:32:00
 * @FilePath: /test/MyConcurrentMap.go
 * @Description:  实现一个并发 map, 要求
 *      1. 支持高并发
 *      2. 仅存在插入和查询操作, 时间复杂度为 O(1)
 *      3. 查询时:
 *         若 key 存在则直接返回 value;
 *         若 key 不存在则阻塞到 key, value 被放入后, 获取 value 返回, 等待指定时长仍未返回则返回超时错误
 *      4. 请使用代码实现, 不能有死锁或 panic 风险
 */
package tutorial

import (
	"context"
	"sync"
	"time"
)

// solution 2
type MyChan struct {
	// 保证 channel 只会被关闭一次
	sync.Once
	ch chan struct{}
}

func NewMyChan() *MyChan {
	return &MyChan{
		ch: make(chan struct{}),
	}
}
func (m *MyChan) Close() {
	m.Do(func() {
		close(m.ch)
	})
}

// 实现一个并发 map, 要求
//  1. 支持高并发
//  2. 仅存在插入和查询操作, 时间复杂度为 O(1)
//  3. 查询时:
//     若 key 存在则直接返回 value;
//     若 key 不存在则阻塞到 key, value 被放入后, 获取 value 返回, 等待指定时长仍未返回则返回超时错误
//  4. 请使用代码实现, 不能有死锁或 panic 风险
type ConcurrentMap struct {
	sync.Mutex
	mp map[int]int
	// key 与相关 channel 的映射关系
	// keyToChan map[int]chan struct{}
	keyToChan map[int]*MyChan
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		mp:        make(map[int]int),
		keyToChan: make(map[int]*MyChan),
	}
}

func (m *ConcurrentMap) Put(k, v int) {
	m.Lock()
	defer m.Unlock()

	ch, ok := m.keyToChan[k]
	if !ok {
		return
	}

	ch.Close()
	// solution 1
	// select {
	// // channel 已关闭直接返回
	// case <-ch:
	// 	return
	// // channel 未关闭,直接关闭
	// default:
	// 	close(ch)
	// }

	// 关闭 channel 以便唤醒所有等待读取的 goroutine
	// 注意不能直接向 ch 放入数据, 因为这样仅会唤醒一个 goroutine
	// close(ch)
}

func (m *ConcurrentMap) Get(key int, maxWaiting time.Duration) (int, error) {
	m.Lock()
	v, ok := m.mp[key]
	if ok {
		m.Unlock()
		return v, nil
	}

	ch, ok := m.keyToChan[key]
	if !ok {
		ch = NewMyChan()
		m.keyToChan[key] = ch
	}

	ctx, cancel := context.WithTimeout(context.Background(), maxWaiting)
	defer cancel()

	// 先解锁再挂起
	m.Unlock()

	select {
	// 超时返回
	case <-ctx.Done():
		return -1, ctx.Err()
	// 已有key数据则返回
	case <-ch.ch:
		break
	}

	// 读取时仍然需要加锁, 防止读的过程中并发写入导致数据不一致
	m.Lock()
	v = m.mp[key]
	m.Unlock()
	return v, nil
}
