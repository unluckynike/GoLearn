package gee

import (
	"net/http"
	"strings"
)

// 之前用了一个非常简单的map结构存储了路由表，使用map存储键值对，索引非常高效，
// 但是有一个弊端，键值对的存储的方式，只能用来索引静态路由。

// 实现动态路由最常用的数据结构，被称为前缀树(Trie树)。
// 每一个节点的所有的子节点都拥有相同的前缀。这种结构非常适用于路由匹配，
type router struct {
	//roots字段是一个字符串到node指针的映射，用于存储不同HTTP请求路径的根节点。
	//每个路径对应一个根节点，根节点包含了该路径下的所有子节点。
	roots map[string]*node
	//handlers字段是一个字符串到HandlerFunc类型的映射，用于存储不同HTTP请求路径对应的处理函数。
	handlers map[string]HandlerFunc
}

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// Only one * is allowed
func parsePattern(pattern string) []string {
	//按照左斜杠切分
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	//并返回一个字符串切片。
	return parts
}

// 在添加路由时，它会将URL模式解析为分段，并将处理函数与该路由的键（由请求方法和URL模式组成）关联起来。
// 然后，它将该路由插入到路由树中，并将处理函数添加到处理程序映射表中。
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *router) handle(c *Context) {
	//在调用匹配到的handler前，将解析出来的路由参数赋值给了c.Params。这样就能够在handler中，通过Context对象访问到具体的值了。
	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		key := c.Method + "-" + n.pattern
		c.Params = params
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	//handle 函数中，将从路由匹配得到的 Handler 添加到 c.handlers列表中，执行c.Next()。
	c.Next()
}
