package gee

import (
	"log"
	"net/http"
	"strings"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
// 定义两个类型： RouterGroup 和 Engine 。
type (
	RouterGroup struct {
		//RouterGroup 类型表示路由组，包含以下字段：
		//-  prefix ：路由组的前缀
		//-  middlewares ：中间件函数的切片，用于支持中间件
		//-  parent ：父路由组，用于支持嵌套
		//-  engine ：所有路由组共享的引擎实例
		prefix      string
		middlewares []HandlerFunc // support middleware
		parent      *RouterGroup  // support nesting
		engine      *Engine       // all groups share a Engine instance
	}

	Engine struct {
		// Engine 类型表示引擎，包含以下字段：
		//-  RouterGroup ：继承自 RouterGroup 类型，表示引擎的默认路由组
		//-  router ：路由器实例，用于处理HTTP请求
		//-  groups ：存储所有路由组的切片
		*RouterGroup
		router *router
		groups []*RouterGroup // store all groups
	}
)

// New is the constructor of gee.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group  将和路由有关的函数，都交给RouterGroup实现
// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// Use is defined to add middleware to the group
// User 方法使用中间件 把他加到某个Group组里
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// 当我们接收到一个具体请求时，要判断该请求适用于哪些中间件，在这里通过 URL 的前缀来判断。得到中间件列表后，赋值给 c.handlers。
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)

	c.handlers = middlewares
	//handle 函数中，将从路由匹配得到的 Handler 添加到 c.handlers列表中，执行c.Next()。
	engine.router.handle(c)
}
