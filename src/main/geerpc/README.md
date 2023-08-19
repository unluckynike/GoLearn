# GoLearn

设计实现go-RPC框架 [GeeRPC](https://geektutu.com/post/geerpc.html)

RPC(Remote Procedure Call，远程过程调用)是一种计算机通信协议，允许调用不同进程空间的程序。RPC 的客户端和服务器可以在一台机器上，也可以在不同的机器上。程序员使用时，就像调用本地程序一样，无需关注内部的实现细节。

Go 语言广泛地应用于云计算和微服务，成熟的 RPC 框架和微服务框架汗牛充栋。grpc、rpcx、go-micro 等都是非常成熟的框架。一般而言，RPC 是微服务框架的一个子集，微服务框架可以自己实现 RPC 部分，当然，也可以选择不同的 RPC 框架作为通信基座。

GeeRPC 选择从零实现 Go 语言官方的标准库 net/rpc，并在此基础上，新增了协议交换(protocol exchange)、注册中心(registry)、服务发现(service discovery)、负载均衡(load balance)、超时处理(timeout processing)等特性。
# ch

- ch01 codec
  - 使用 encoding/gob 实现消息的编解码(序列化与反序列化)
  - 实现一个简易的服务端，仅接受消息，不处理
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