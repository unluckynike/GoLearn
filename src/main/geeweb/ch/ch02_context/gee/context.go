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
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
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

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
