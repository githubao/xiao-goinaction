// 主函数
// author: baoqiang
// time: 2018/6/2 下午3:46
package main

import (
	"log"
	"os"
	"github.com/githubao/xiao-goinaction/chapter2/sample/search"
)

func init()  {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
