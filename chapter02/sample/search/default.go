// 默认的matcher
// author: baoqiang
// time: 2018/6/4 上午11:01
package search

type defaultMatcher struct {

}

func init()  {
	var matcher defaultMatcher
	Register("default",matcher)
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([] *Result, error) {
	return nil,nil
}