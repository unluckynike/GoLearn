# GoLearn

设计实现go-ORM框架 [GeeORM](https://geektutu.com/post/geeorm.html)

对象关系映射（Object Relational Mapping，简称ORM）是通过使用描述对象和数据库之间映射的元数据，将面向对象语言程序中的对象自动持久化到关系数据库中。

ORM 框架相当于对象和数据库中间的一个桥梁，借助 ORM 可以避免写繁琐的 SQL 语言，仅仅通过操作具体的对象，就能够完成对关系型数据库的操作。

Go 语言中使用比较广泛 ORM 框架是 gorm 和 xorm。除了基础的功能，比如表的操作，记录的增删查改，gorm 还实现了关联关系(一对一、一对多等)，回调插件等；xorm 实现了读写分离(支持配置多个数据库)，数据同步，导入导出等。

如何根据任意类型的指针，得到其对应的结构体的信息。这涉及到了 Go 语言的反射机制(reflect)，通过反射，可以获取到对象对应的结构体名称，成员变量、方法等信息，

| 数据库概念       | 面向对象编程语言对应概念  |
|-----------------|------------------------|
| 表 (table)       | 类 (class/struct)      |
| 记录 (record, row) | 对象 (object)         |
| 字段 (field, column) | 对象属性 (attribute) |

# 数据库

[SQLite](https://www.sqlite.org/download.html) 

[window 安装 Sqlite 并连接](https://blog.csdn.net/qq_30718137/article/details/117882063?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522169244813316800226593280%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=169244813316800226593280&biz_id=0&utm_medium=distribute.pc_chrome_plugin_search_result.none-task-blog-2~all~first_rank_ecpm_v1~rank_v31_ecpm-3-117882063-null-null.nonecase&utm_term=windos%20%E5%AE%89%E8%A3%85SQLlite&spm=1018.2226.3001.4187)


# ch

- ch01 database-sql
  - SQLite 的基础操作（连接数据库，创建表、增删记录等）。
  - 使用 Go 语言标准库 database/sql 连接并操作 SQLite 数据库，并简单封装
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