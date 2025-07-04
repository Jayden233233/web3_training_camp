// 指针
// ✅题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。

// ✅题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。

package main

import (
	"fmt"
)

func main() {
	num := 5         // 初始化一个整数变量
	increment(&num)  // 传递变量地址给函数
	fmt.Println(num) // 输出修改后的值

	slice := []int{1, 2, 3}
	fmt.Println(slice)
	multi2(&slice)
	fmt.Println(slice)
}

func increment(a *int) {
	*a += 10
}

func multi2(slicePtr *[]int) {
	for i, _ := range *slicePtr {
		(*slicePtr)[i] = 2 * (*slicePtr)[i]

	}
}
