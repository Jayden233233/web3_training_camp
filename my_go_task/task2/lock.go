// 锁机制
// ✅题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

// ✅题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	firstImpl()
	secondImpl()

}

func firstImpl() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	var count int = 0
	wg.Add(10)

	for j := 0; j < 10; j++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				lock.Lock()
				count += 1
				lock.Unlock()
			}

		}()
	}
	wg.Wait()                          // 等待所有 goroutine 完成
	fmt.Println("Final count:", count) // 输出最终结果
}

func secondImpl() {
	var wg sync.WaitGroup
	var count int32 = 0
	wg.Add(10)

	for j := 0; j < 10; j++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt32(&count, 1) // 原子递增操作
			}

		}()
	}
	wg.Wait()                          // 等待所有 goroutine 完成
	fmt.Println("Final count:", count) // 输出最终结果
}
