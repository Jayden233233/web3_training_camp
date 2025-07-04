// 面向对象
// ✅题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。

// ✅题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。
package main

import "fmt"

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func (this *Rectangle) Area() {
	fmt.Println("this is Area of Rectangle")

}

func (this *Rectangle) Perimeter() {
	fmt.Println("this is Perimeter of Rectangle")
}

type Circle struct {
}

func (this *Circle) Area() {
	fmt.Println("this is Area of Circle")
}

func (this *Circle) Perimeter() {
	fmt.Println("this is Perimeter of Circle")
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (this *Employee) PrintInfo() {
	fmt.Println(this.Name)
	fmt.Println(this.Age)
	fmt.Println(this.EmployeeID)

}
func main() {
	r := Rectangle{}
	r.Area()
	r.Perimeter()

	c := Circle{}
	c.Area()
	c.Perimeter()

	d := Employee{
		Person: Person{
			Name: "asc",
			Age:  3,
		},

		EmployeeID: 789,
	}
	d.PrintInfo()

}
