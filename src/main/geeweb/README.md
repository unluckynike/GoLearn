# GoLearn

设计go-web框架

[Gee](https://geektutu.com/post/gee.html)

# ch

- ch01 http
  - 简单介绍net/http库以及http.Handler接口。
  - 搭建Gee框架的雏形。
- ch02 上下文
    - 将路由(router)独立出来，方便之后增强。
    - 设计上下文(Context)，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持。
- ch03 前缀树路由
    - 使用 Trie 树实现动态路由(dynamic route)解析。
    - 支持两种模式:name和*filepath。
- ch04 分组控制
    -  实现路由分组控制(Route Group Control)。
- ch05 中间件
    - 设计并实现 Web 框架的中间件(Middlewares)机制。
    - 实现通用的Logger中间件，能够记录请求到响应所花费的时间。
- ch06 模板Template
    - 实现静态资源服务(Static Resource)。
    - 支持HTML模板渲染。  
- ch07 错误恢复
    - 实现错误处理机制。