package gee

type node struct {
	pattern  string  // 待匹配路由 比如/p/:lang
	part     string  // 路由中的一部分 比如:lang
	children []*node // 子节点，比如[doc,intro]
	isWild   bool    //是否精确匹配 part含有：或* 时为true
}
