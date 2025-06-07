很好！Go（Golang）凭借轻量级的 Goroutine 和原生调度器支持，是模拟**抢占式调度**和**上下文切换**的理想选择。Go 的运行时已经内置了抢占式调度（基于时间片和系统调用），我们可以直接利用 Goroutine 和 Channel 来模拟操作系统的行为。

---

## **Go 实现抢占式调度与上下文切换**

### **1. 抢占式调度模拟**
Go 的调度器本身是抢占式的（从 Go 1.14 开始支持**非协作式抢占**），但我们可以手动模拟更细粒度的任务切换，比如：
- **基于时间片的抢占**（固定时间强制切换任务）
- **基于优先级的抢占**（高优先级任务抢占低优先级任务）
- **模拟保存/恢复上下文**（类似操作系统切换进程）

#### **示例：时间片轮转调度（Round-Robin）**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

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
}
```
**运行结果：**
```
Task 1: Running...
Task 2: Running...
Task 3: Running...
Task 1: Running...
Task 2: Running...
Task 3: Running...
...（持续 3 秒）
Task 1: Stopped by scheduler
Task 2: Stopped by scheduler
Task 3: Stopped by scheduler
All tasks stopped by scheduler.
```

**关键点：**
- `stopChan` 用于通知 Goroutine 停止（模拟抢占）。
- `time.Sleep(500ms)` 模拟任务执行时间片。
- 主 Goroutine 在 3 秒后强制关闭所有任务（模拟调度器中断）。

---

### **2. 模拟上下文切换**
操作系统的上下文切换（Context Switch）需要保存/恢复寄存器状态，Go 的 Goroutine 切换由调度器自动管理，但我们可以模拟类似行为：
- **手动保存任务状态**（如 `struct Task`）
- **模拟调度器选择下一个任务**

#### **示例：手动任务切换（模拟上下文保存）**
```go
package main

import (
	"fmt"
	"time"
)

// 模拟任务上下文
type Task struct {
	ID     int
	State  string // "ready", "running", "blocked"
	PC     int    // 模拟程序计数器
}

func main() {
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
```
**运行结果：**
```
Scheduler: Switching to Task 1 (PC=0)
Scheduler: Switching to Task 2 (PC=0)
Scheduler: Switching to Task 3 (PC=0)
Scheduler: Switching to Task 1 (PC=1)
Scheduler: Switching to Task 2 (PC=1)
Scheduler: Switching to Task 3 (PC=1)
...
```
**关键点：**
- `Task` 结构体模拟任务上下文（PC 寄存器、状态）。
- 调度器手动选择下一个任务（模拟时间片轮转）。
- `time.Sleep(300ms)` 模拟任务执行时间。

---

## **3. 更高级的模拟（优先级抢占）**
如果想模拟**优先级调度**，可以：
1. 为任务分配优先级（`Priority int`）。
2. 调度器选择最高优先级的任务执行。
3. 高优先级任务可以抢占低优先级任务。

#### **示例：优先级抢占**
```go
type Task struct {
	ID       int
	Priority int // 优先级（1=最高）
	State    string
}

func main() {
	tasks := []*Task{
		{ID: 1, Priority: 2, State: "ready"},
		{ID: 2, Priority: 1, State: "ready"}, // 最高优先级
		{ID: 3, Priority: 3, State: "ready"},
	}

	// 调度器选择最高优先级任务
	for i := 0; i < 5; i++ {
		var highest *Task
		for _, t := range tasks {
			if highest == nil || t.Priority < highest.Priority {
				highest = t
			}
		}

		fmt.Printf("Scheduler: Running Task %d (Priority=%d)\n", highest.ID, highest.Priority)
		time.Sleep(500 * time.Millisecond)
	}
}
```
**运行结果：**
```
Scheduler: Running Task 2 (Priority=1)
Scheduler: Running Task 2 (Priority=1)
Scheduler: Running Task 2 (Priority=1)
...（Task 2 一直抢占，直到完成）
```

---

## **总结**
| 功能           | Go 实现方式                       |
| -------------- | --------------------------------- |
| **抢占式调度** | `channel` + `time.Sleep` 模拟中断 |
| **上下文切换** | `struct Task` 保存 PC/状态        |
| **优先级调度** | 遍历选择最高优先级任务            |

Go 的 Goroutine 本身是抢占式的，但我们可以手动模拟更细粒度的调度策略，适合学习操作系统调度原理！ 🚀