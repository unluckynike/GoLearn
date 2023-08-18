package gee

import (
	"fmt"
	"log"
	"net/http"
)

/*
实现了路由映射表，提供了用户注册静态路由的方法，包装了启动服务的函数。
*/

// HandlerFunc 一个类型为HandlerFunc的函数类型
// 提供给框架用户的，用来定义路由映射的处理方法
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	//一张路由映射表router，key 由请求方法和静态路由地址构成，
	//例如GET-/、GET-/hello、POST-/hello，这样针对相同的路由，
	//请求方法不同,可以映射不同的处理方法(Handler)，value 是用户映射的处理方法。
	router map[string]HandlerFunc
}

// New is the constructor of gee.Engine
// 创建了一个新的Engine实例，其中包含一个空的映射用于存储路由处理函数的映射关系。
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 将路由信息添加到引擎的路由表中。
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handler
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动http sever 像极了Gin
// 用户调用(*Engine).GET()方法时，会将路由和处理方法注册到映射表 router 中，
// (*Engine).Run()方法，是 ListenAndServe 的包装。
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// Engine实现的 ServeHTTP 方法的作用就是，解析请求的路径，查找路由映射表，
// 如果查到，就执行注册的处理方法。如果查不到，就返回 404 NOT FOUND
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
