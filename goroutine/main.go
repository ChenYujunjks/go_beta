package main

import (
	"fmt"
	"sync"
	"time"
)

// 模拟任务上下文
type Task struct {
	ID     int
	State  string // "ready", "running", "blocked"
	PC     int    // 模拟程序计数器
}

func task(id int, wg *sync.WaitGroup, stopChan chan struct{}) {
	defer wg.Done()

	for {
		select {
		case <-stopChan:
			fmt.Printf("Task %d: Stopped by scheduler\n", id)
			return
		default:
			fmt.Printf("Task %d: Running...\n", id)
			time.Sleep(500 * time.Millisecond) // 模拟任务执行
		}
	}
}

func main() {
	//Round Robin
	var wg sync.WaitGroup
	stopChan := make(chan struct{})

	// 启动 3 个任务（Goroutine）
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go task(i, &wg, stopChan)
	}

	// 模拟调度器：运行 3 秒后强制停止所有任务
	time.Sleep(3 * time.Second)
	close(stopChan) // 发送停止信号（模拟抢占）

	wg.Wait()
	fmt.Println("All tasks stopped by scheduler.")
	Context_switch()
}


func Context_switch() {
	tasks := []*Task{
		{ID: 1, State: "ready", PC: 0},
		{ID: 2, State: "ready", PC: 0},
		{ID: 3, State: "ready", PC: 0},
	}

	currentTask := -1

	// 模拟调度器
	for i := 0; i < 10; i++ {
		// 选择下一个任务（轮询）
		currentTask = (currentTask + 1) % len(tasks)
		task := tasks[currentTask]

		// 模拟上下文切换
		fmt.Printf("Scheduler: Switching to Task %d (PC=%d)\n", task.ID, task.PC)
		task.State = "running"

		// 模拟任务执行（修改 PC）
		task.PC += 1
		time.Sleep(300 * time.Millisecond) // 模拟执行时间片

		// 模拟任务返回（状态变回 ready）
		task.State = "ready"
	}
}