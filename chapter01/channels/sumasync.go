// 使用管道计算加法
// author: baoqiang
// time: 2018/7/9 下午7:59
package main

import "fmt"

func main() {
	c1 := NewCalculator()
	sum := c1.Sum(2, 3)
	fmt.Println(sum)
}

type Cmd struct {
	A        int
	B        int
	Sum      int
	doneChan chan *Cmd
}

type Calculator struct {
	cQueue chan *Cmd
}

func NewCalculator() *Calculator {
	c1 := &Calculator{cQueue: make(chan *Cmd, 10)}
	go c1.calculate()
	return c1
}

func (c1 *Calculator) Sum(a, b int) int {
	fmt.Println("run sum, now we use async sum")
	c := c1.sumAsync(a, b)
	res := <- c.doneChan
	return res.Sum
}

func (c1 *Calculator) sumAsync(a, b int) *Cmd {
	fmt.Println("construct a new instance push into queqe")
	c := &Cmd{A: a, B: b, doneChan: make(chan *Cmd)}
	c1.cQueue <- c
	return c
}

func (c1 *Calculator) calculate() {
	fmt.Println("waiting data for summing...")
	for c := range c1.cQueue {
		c.Sum = c.A + c.B
		c.doneChan <- c
	}
}
