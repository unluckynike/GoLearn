package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 创建自定义类型H，是一个映射类型(给map[string]interface{}起了一个别名gee.H，构建JSON数据时，显得更简洁。)
// 可以用于存储:
// 键 为字符串类型
// 值 为任意类型的数据。
type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	// request info
	Req    *http.Request
	Path   string
	Method string
	Params map[string]string //路由参数访问
	// 响应信息
	StatusCode int
	// middleware 中间件
	handlers []HandlerFunc
	index    int //index是记录当前执行到第几个中间件
	// engine pointer
	//添加了成员变量 engine *Engine，这样就能够通过 Context 访问 Engine 中的 HTML 模板。实例化 Context 时，还需要给 c.engine 赋值。
	engine *Engine
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

// index是记录当前执行到第几个中间件，当在中间件中调用Next方法时，控制权交给了下一个中间件，直到调用到最后一个中间件，
// 然后再从后往前，调用每个中间件在Next方法之后定义的部分
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}
func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

// 提供对路由参数的访问。 解析后的参数存储到Params中
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// 提供了访问Query和PostForm参数的方法。
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// 得到URL中的字段 ?name=xxx
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 设置状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 设置请求头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 提供了快速构造String/Data/JSON/HTML响应的方法。
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}
