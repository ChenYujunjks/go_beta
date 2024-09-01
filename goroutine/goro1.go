package goroutine

import (
	"fmt"
	"time"
)

func printNumbers(c chan int) {
	for number := range c {
		fmt.Println(number)
		time.Sleep(1 * time.Second)
	}
}

func DemoConcurrency() {
	c := make(chan int)
	go printNumbers(c)
	for i := 1; i <= 3; i++ {
		c <- i
	}
	close(c)
}
