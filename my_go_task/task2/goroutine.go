// Goroutine
// ✅题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	start_two_go_routine(&wg)
	wg.Wait()
}

func start_two_go_routine(wg *sync.WaitGroup) {
	go print_odd(wg)
	go print_even(wg)
}

func print_odd(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}

func print_even(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
