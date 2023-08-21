package gee

import (
	"html/template"
	"log"
	"net/http"
	"path"
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
		//Go语言内置了text/template和html/template2个模板标准库，
		//其中html/template为 HTML 提供了较为完整的支持。包括普通变量渲染、列表渲染、对象渲染等。
		//*template.Template 和 template.FuncMap对象，前者将所有的模板加载进内存，后者是所有的自定义模板渲染函数。
		htmlTemplates *template.Template // for html render
		funcMap       template.FuncMap   // for html render
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

// 解析请求的地址，映射到服务器上文件的真实地址，交给http.FileServer处理。
// create static handler
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// Check if files exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// serve static files
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	group.GET(urlPattern, handler)
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// 给用户分别提供了设置自定义渲染函数funcMap和加载模板的方法。
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}
func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
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
	c.engine = engine
	c.handlers = middlewares
	//handle 函数中，将从路由匹配得到的 Handler 添加到 c.handlers列表中，执行c.Next()。
	engine.router.handle(c)
}
