// ✅题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 表示一个任务，是一个无参数、无返回值的函数
type Task func()

// TaskResult 存储任务的执行结果和执行时间
type TaskResult struct {
	TaskID   int
	Duration time.Duration
	Err      error
}

// TaskScheduler 任务调度器
type TaskScheduler struct {
	maxWorkers int
	tasks      []Task
	results    []TaskResult
	wg         sync.WaitGroup
	taskQueue  chan Task
	resultChan chan TaskResult
	mu         sync.Mutex
}

// NewTaskScheduler 创建一个新的任务调度器
func NewTaskScheduler(maxWorkers int) *TaskScheduler {
	return &TaskScheduler{
		maxWorkers: maxWorkers,
		tasks:      make([]Task, 0),
		results:    make([]TaskResult, 0),
		taskQueue:  make(chan Task),
		resultChan: make(chan TaskResult),
	}
}

// AddTask 添加任务到调度器
func (ts *TaskScheduler) AddTask(task Task) {
	ts.tasks = append(ts.tasks, task)
}

// Start 启动任务调度器
func (ts *TaskScheduler) Start() {
	// 启动工作协程
	for i := 0; i < ts.maxWorkers; i++ {
		ts.wg.Add(1)
		go ts.worker(i)
	}

	// 启动结果收集协程
	go ts.collectResults()

	// 将任务发送到任务队列
	go func() {
		for i, task := range ts.tasks {
			taskID := i
			ts.taskQueue <- func() {
				ts.executeTask(taskID, task)
			}
		}
		close(ts.taskQueue)
	}()

	// 等待所有工作协程完成
	ts.wg.Wait()
	close(ts.resultChan)
}

// worker 工作协程，处理任务
func (ts *TaskScheduler) worker(id int) {
	defer ts.wg.Done()
	for task := range ts.taskQueue {
		task()
	}
}

// executeTask 执行单个任务并记录执行时间
func (ts *TaskScheduler) executeTask(taskID int, task Task) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		var err error
		if r := recover(); r != nil {
			err = fmt.Errorf("task panicked: %v", r)
		}
		ts.resultChan <- TaskResult{
			TaskID:   taskID,
			Duration: duration,
			Err:      err,
		}
	}()
	task()
}

// collectResults 收集任务执行结果
func (ts *TaskScheduler) collectResults() {
	for result := range ts.resultChan {
		ts.mu.Lock()
		ts.results = append(ts.results, result)
		ts.mu.Unlock()
	}
}

// GetResults 获取所有任务的执行结果
func (ts *TaskScheduler) GetResults() []TaskResult {
	return ts.results
}

func main() {
	// 创建一个最大并发数为3的任务调度器
	scheduler := NewTaskScheduler(3)

	// 添加一些测试任务
	for i := 0; i < 5; i++ {
		taskID := i
		scheduler.AddTask(func() {
			// 模拟不同任务的执行时间
			duration := time.Duration(taskID+1) * time.Second
			time.Sleep(duration)
			fmt.Printf("Task %d completed after %v\n", taskID, duration)
		})
	}

	// 启动调度器
	scheduler.Start()

	// 打印所有任务的执行时间
	fmt.Println("\n任务执行统计：")
	for _, result := range scheduler.GetResults() {
		if result.Err != nil {
			fmt.Printf("Task %d failed: %v, duration: %v\n", result.TaskID, result.Err, result.Duration)
		} else {
			fmt.Printf("Task %d executed in %v\n", result.TaskID, result.Duration)
		}
	}
}
