// main
// author: baoqiang
// time: 2018/6/3 下午4:56
package main

import (
	"fmt"
	"time"
)

func main1() {
	sli := make([]int, 1, 3)

	//sli[0] = 2
	sli[1] = 2

	fmt.Print(sli)
}

func main() {
	current := time.Now().Format("2006-01-02 15:04:05")
	text := fmt.Sprintf("i recieved %s at %s", "aaa", current)
	fmt.Print(text)
}
