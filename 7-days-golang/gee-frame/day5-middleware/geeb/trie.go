package geeb

import (
	"fmt"
	"strings"
)

// Trie 树实现
type node struct {
	pattern  string  // 待匹配路由，例如：/p/:lang
	part     string  // 路由中的一部分，例如：:lang
	children []*node // 子节点，例如：[doc,tutorial,intro]
	isWild   bool    // 是否精准匹配，part 含有 ：或者 * 时为true
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s,part=%s,isWild=%t}", n.pattern, n.part, n.isWild)
}

// 对于路由来说，最重要的当然是注册与匹配了。
// 开发服务时，注册路由规则，映射handler；访问时，匹配路由规则，查找对应的handler，因此，Trie树需要支持节点的插入与查询。
// 插入功能很简单，递归查找每一层的节点，如果没有匹配到当前part的节点，则新建一个，有一点需要注意，/p/:lang/doc只有在第
// 三层节点，即doc节点，pattern才会设置为/p/:lang/doc。p和:lang节点的pattern属性皆为空。因此，当匹配结束时，我们可以
// 使用n.pattern == ""来判断路由规则是否匹配成功。例如，/p/python虽能成功匹配到:lang，但:lang的pattern值为空，因此
// 匹配失败。
// 查询功能，同样也是递归查询每一层的节点，退出规则是，匹配到了*，匹配失败，或者匹配到了第len(parts)层节点。

// 插入-递归
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 查找-递归
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

// 与普通的树不同，为了实现动态路由匹配，加上了isWild这个参数。即当我们匹配 /p/go/doc这个路由时，
// 第一层节点，p精准匹配到了p，第二层节点，go模糊匹配到 :lang，那么会把lang这个参数赋值为go，继
// 续下一层匹配。我们将匹配的逻辑，包装为一个辅助函数。

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用户查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
