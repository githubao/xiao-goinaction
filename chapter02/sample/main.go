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
	// 2018/06/04 11:45:16 Description:
	//The January letter asserts the president's power to end the special counsel probe and to issue pardons, and reveals Trump's broad interpretation of executive power.
	search.Run("president")
}
