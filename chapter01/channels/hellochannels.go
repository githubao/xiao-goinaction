// go routine demo
// author: baoqiang
// time: 2018/6/4 上午11:48
package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("Recieved: %d\n", i)
	}

	wg.Done()
}

func main() {
	c := make(chan int)

	go printer(c)

	wg.Add(1)

	for i := 1; i <= 10; i++ {
		c <- i
	}

	close(c)

	wg.Wait()
}
