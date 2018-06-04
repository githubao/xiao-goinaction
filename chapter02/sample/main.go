// 主函数
// author: baoqiang
// time: 2018/6/2 下午3:46
package main

import (
	"log"
	"os"

	//"github.com/githubao/xiao-goinaction/chapter02/sample/search/..."
	"xiao-goinaction/chapter02/sample/search"
	_ "xiao-goinaction/chapter02/sample/matchers"
)

func init()  {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
