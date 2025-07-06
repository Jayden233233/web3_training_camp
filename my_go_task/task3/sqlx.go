// Sqlx入门
// 题目1：使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
// 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Employee 结构体映射employees表
type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

// Book 结构体映射books表
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	// 连接SQLite内存数据库（无需文件）
	db, err := sqlx.Connect("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	defer db.Close()

	// 创建表
	if err := createTables(db); err != nil {
		log.Fatalf("创建表失败: %v", err)
	}

	// 插入测试数据
	if err := seedData(db); err != nil {
		log.Fatalf("插入数据失败: %v", err)
	}

	// 题目1-1：查询技术部员工
	techEmployees, err := getEmployeesByDepartment(db, "技术部")
	if err != nil {
		log.Fatalf("查询技术部员工失败: %v", err)
	}
	fmt.Println("技术部员工:")
	for _, emp := range techEmployees {
		fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %d\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	// 题目1-2：查询工资最高的员工
	topEmployee, err := getHighestSalaryEmployee(db)
	if err != nil {
		log.Fatalf("查询最高工资员工失败: %v", err)
	}
	fmt.Printf("\n工资最高的员工: %s (工资: %d)\n", topEmployee.Name, topEmployee.Salary)

	// 题目2：查询价格大于50的书籍
	expensiveBooks, err := getBooksOverPrice(db, 50.0)
	if err != nil {
		log.Fatalf("查询高价书籍失败: %v", err)
	}
	fmt.Println("\n价格大于50的书籍:")
	for _, book := range expensiveBooks {
		fmt.Printf("ID: %d, 书名: %s, 作者: %s, 价格: %.2f\n", book.ID, book.Title, book.Author, book.Price)
	}
}

// 创建测试表
func createTables(db *sqlx.DB) error {
	_, err := db.Exec(`
		CREATE TABLE employees (
			id INTEGER PRIMARY KEY,
			name TEXT,
			department TEXT,
			salary INTEGER
		);

		CREATE TABLE books (
			id INTEGER PRIMARY KEY,
			title TEXT,
			author TEXT,
			price REAL
		);
	`)
	return err
}

// 插入测试数据
func seedData(db *sqlx.DB) error {
	// 插入员工数据
	employees := []struct {
		Name       string
		Department string
		Salary     int
	}{
		{"张三", "技术部", 10000},
		{"李四", "市场部", 8000},
		{"王五", "技术部", 12000},
		{"赵六", "财务部", 9000},
	}

	for _, emp := range employees {
		_, err := db.Exec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)",
			emp.Name, emp.Department, emp.Salary)
		if err != nil {
			return err
		}
	}

	// 插入书籍数据
	books := []struct {
		Title  string
		Author string
		Price  float64
	}{
		{"Go语言编程", "Alan A. Donovan", 89.0},
		{"SQL必知必会", "Ben Forta", 45.0},
		{"数据结构与算法", "Thomas H. Cormen", 128.0},
		{"Python基础教程", "Magnus Lie Hetland", 68.0},
	}

	for _, book := range books {
		_, err := db.Exec("INSERT INTO books (title, author, price) VALUES (?, ?, ?)",
			book.Title, book.Author, book.Price)
		if err != nil {
			return err
		}
	}

	return nil
}

// 题目1-1：查询指定部门的员工
func getEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM employees WHERE department = ?", department)
	return employees, err
}

// 题目1-2：查询工资最高的员工
func getHighestSalaryEmployee(db *sqlx.DB) (Employee, error) {
	var employee Employee
	err := db.Get(&employee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	return employee, err
}

// 题目2：查询价格大于指定值的书籍
func getBooksOverPrice(db *sqlx.DB, price float64) ([]Book, error) {
	var books []Book
	err := db.Select(&books, "SELECT * FROM books WHERE price > ?", price)
	return books, err
}
