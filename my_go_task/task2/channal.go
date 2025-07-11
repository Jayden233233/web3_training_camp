// Channel
// ✅题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。

// ✅题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var ch = make(chan int, 10)

	go write(ch, &wg)
	go read(ch, &wg)

	wg.Wait()

}

func write(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 通知 WaitGroup 读取完成
	for i := 1; i <= 20; i++ {
		ch <- i
	}
	close(ch) // 关闭 channel 表示不再有数据发送

}

func read(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 通知 WaitGroup 读取完成
	for a := range ch {
		fmt.Println(a)
	}

}
