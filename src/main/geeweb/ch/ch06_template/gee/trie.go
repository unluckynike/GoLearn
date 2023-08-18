package gee

import "strings"

// 动态路由前缀树(Trie树)。
type node struct {
	//设计数节点上应该存储的信息
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
// matchChild 返回与给定部分匹配的子节点或第一个通配符子节点。 如果没有匹配的子节点，则返回nil。
// 参数：part - 要匹配的部分字符串
// 返回值：匹配的子节点或nil
// 注意：isWild属性表示子节点是否为通配符节点。如果子节点的part属性与part相同，或者子节点的isWild属性为true，则表示匹配成功。
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 开发服务时，注册路由规则，映射handler；访问时，匹配路由规则，查找到对应的handler。
// 因此，Trie 树需要支持节点的插入与查询。
func (n *node) insert(pattern string, parts []string, height int) {

	//递归查找每一层的节点，如果没有匹配到当前part的节点，则新建一个，
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
