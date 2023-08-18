# GoLearn

Go语言中文网 

[Go 系列教程](https://studygolang.com/subject/2)

# 
- Go 系列教程 —— 36. 写入文件
  - ![write files](https://raw.githubusercontent.com/studygolang/gctt-images/master/golang-series/golang-write-files.png) 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 36 篇。 在这一章我们将学习如何使用 Go 语言将数据写到文件里面。并且还要学习如何同步的写到文件里面。 这章教程包括如下几个部分：...


- Go 系列教程 —— 35. 读取文件
  - ![reading files](https://raw.githubusercontent.com/studygolang/gctt-images/master/golang-series/golang-read-files.png) 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 35 篇。 文件读取是所有编程语言中最常见的操作之一。本教程我们会学习如何使用 Go 读取文件。 本教程分为如下小节。 - 将...


- Go 系列教程 —— 34. 反射
  - ![reflection](https://raw.githubusercontent.com/studygolang/gctt-images/master/golang-series/reflection-golang-3.png) 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 34 篇。 反射是 Go 语言的高级主题之一。我会尽可能让它变得简单易懂。 本教程分为如下小节。 - 什么是反射？ - 为何需...


- Go 系列教程 —— 33. 函数是一等公民（头等函数）
  - ![custom errors](https://raw.githubusercontent.com/studygolang/gctt-images/master/golang-series/first-class-functions-golang.png) 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 33 篇。 ## 什么是头等函数？ **支持头等函数（First Class Function）的编程语言，可...


- Go 系列教程 —— 32. panic 和 recover
  - [panic 和 recover](https://raw.githubusercontent.com/studygolang/gctt-images/master/golang-series/panic-recover-golang-2-2.png) 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 32 篇。 ## 什么是 panic？ 在 Go 语言中，程序中一般是使用[错误](https://studygol...


- Go 系列教程 —— 31. 自定义错误
  - ![custom errors](https://raw.githubusercontent.com/studygolang/gctt-images/master/golang-series/custom-errors-golang-1.png) 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 31 篇。 在[上一教程](https://studygolang.com/articles/12724)里，我们学习了 Go ...


- Go 系列教程 —— 30. 错误处理
  - 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 30 篇。 ## 什么是错误？ 错误表示程序中出现了异常情况。比如当我们试图打开一个文件时，文件系统里却并没有这个文件。这就是异常情况，它用一个错误来表示。 在 Go 中，错误一直是很常见的。错误用内建的 `error` 类型来表示。 就像其他的内建类型（如 `int`、`float64` 等），错误值可以存储在变量里、作为函数的返回值等等。 ## 示例 ...


- Go 系列教程 —— 29. Defer
  - 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 29 篇。 ## 什么是 defer？ `defer` 语句的用途是：含有 `defer` 语句的函数，会在该函数将要返回之前，调用另一个函数。这个定义可能看起来很复杂，我们通过一个示例就很容易明白了。 ## 示例 ```go package main import ( "fmt" ) func finished() { fmt....

- Go 系列教程 —— 28. 多态
  - 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 28 篇。 Go 通过[接口](https://studygolang.com/articles/12266)来实现多态。我们已经讨论过，在 Go 语言中，我们是隐式地实现接口。一个类型如果定义了接口所声明的全部[方法](https://studygolang.com/articles/12264)，那它就实现了该接口。现在我们来看看，利用接口，Go 是如何实现多态的。 #...

- Go 系列教程 —— 27. 组合取代继承
  - 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 27 篇。 Go 不支持继承，但它支持组合（Composition）。组合一般定义为“合并在一起”。汽车就是一个关于组合的例子：一辆汽车由车轮、引擎和其他各种部件组合在一起。 ## 通过嵌套结构体进行组合 在 Go 中，通过在结构体内嵌套结构体，可以实现组合。 组合的典型例子就是博客帖子。每一个博客的帖子都有标题、内容和作者信息。使用组合可以很好地表示它们。通过...

- Go 系列教程 —— 26. 结构体取代类
  - 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 26 篇。 ## Go 支持面向对象吗？ Go 并不是完全面向对象的编程语言。Go 官网的 [FAQ](https://golang.org/doc/faq#Is_Go_an_object-oriented_language) 回答了 Go 是否是面向对象语言，摘录如下。 > 可以说是，也可以说不是。虽然 Go 有类型和方法，支持面向对象的编程风格，但却没有类型的层次...

- Go 系列教程 —— 25. Mutex
  - 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 25 篇。 本教程我们学习 Mutex。我们还会学习怎样通过 Mutex 和[信道](https://studygolang.com/articles/12402)来处理竞态条件（Race Condition）。 ## 临界区 在学习 Mutex 之前，我们需要理解并发编程中临界区（Critical Section）的概念。当程序并发地运行时，多个 [Go 协...

- Go 系列教程 —— 24. Select
  - 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 24 篇。 ## 什么是 select？ `select` 语句用于在多个发送/接收信道操作中进行选择。`select` 语句会一直阻塞，直到发送/接收操作准备就绪。如果有多个信道操作准备完毕，`select` 会随机地选取其中之一执行。该语法与 `switch` 类似，所不同的是，这里的每个 `case` 语句都是信道操作。我们好好看一些代码来加深理解吧。 ##...

- Go 系列教程 —— 23. 缓冲信道和工作池（Buffered Channels and Worker Pools）
  - 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 23 篇。 ## 什么是缓冲信道？ 在[上一教程](https://studygolang.com/articles/12402)里，我们讨论的主要是无缓冲信道。我们在[信道](https://studygolang.com/articles/12402)的教程里详细讨论了，无缓冲信道的发送和接收过程是阻塞的。 我们还可以创建一个有缓冲（Buffer）的信道。...

- Go 系列教程 —— 22. 信道（channel）
  - 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 22 篇。 在[上一教程](https://studygolang.com/articles/12342)里，我们探讨了如何使用 Go 协程（Goroutine）来实现并发。我们接着在本教程里学习信道（Channel），学习如何通过信道来实现 Go 协程间的通信。 ## 什么是信道？ 信道可以想像成 Go 协程之间通信的管道。如同管道中的水会从一端流到另一端，...

- Go 系列教程 —— 21. Go 协程
  - 欢迎来到 [Golang 系列教程](https://studygolang.com/subject/2)的第 21 篇。 在前面的教程里，我们探讨了并发，以及并发与并行的区别。本教程则会介绍在 Go 语言里，如何使用 Go 协程（Goroutine）来实现并发。 ## Go 协程是什么？ Go 协程是与其他函数或方法一起并发运行的函数或方法。Go 协程可以看作是轻量级线程。与线程相比，创建一个 Go 协程的成本很小。因此在 Go 应用中，常常会看到有数以千计的 Go 协程...

- Go 系列教程 —— 20. 并发入门
  - 欢迎来到我们 [Golang 系列教程](https://studygolang.com/subject/2)的第 20 篇。 **Go 是并发式语言，而不是并行式语言**。在讨论 Go 如何处理并发之前，我们必须理解何为并发，以及并发与并行的区别。 ## 并发是什么？ 并发是指立即处理多个任务的能力。一个例子就能很好地说明这一点。 我们可以想象一个人正在跑步。假如在他晨跑时，鞋带突然松了。于是他停下来，系一下鞋带，接下来继续跑。这个例子就是典型的并发。这个人...

- Go 系列教程 —— 19. 接口（二）
  - 欢迎来到 Golang 系列教程的第 19 个教程。接口共有两个教程，这是我们第二个教程。如果你还没有阅读前面的教程，请你阅读[接口（一）](https://studygolang.com/articles/12266)。 ### 实现接口：指针接受者与值接受者 在[接口（一）](https://studygolang.com/articles/12266)上的所有示例中，我们都是使用值接受者（Value Receiver）来实现接口的。我们同样可以使用指针接受者（Pointer R...

- Go 系列教程 —— 18. 接口（一）
  - 欢迎来到 [Golang 系列教程](/subject/2)的第 18 个教程。接口共有两个教程，这是我们接口的第一个教程。 ### 什么是接口？ 在面向对象的领域里，接口一般这样定义：**接口定义一个对象的行为**。接口只指定了对象应该做什么，至于如何实现这个行为（即实现细节），则由对象本身去确定。 在 Go 语言中，接口就是方法签名（Method Signature）的集合。当一个类型定义了接口中的所有方法，我们称它实现了该接口。这与面向对象编程（OOP）的说法很类...

- Go 系列教程 —— 17. 方法
  - 欢迎来到 [Golang 系列教程](/subject/2) 的第 17 个教程。 ### 什么是方法？ 方法其实就是一个函数，在 `func` 这个关键字和方法名中间加入了一个特殊的接收器类型。接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的。 下面就是创建一个方法的语法。 ```go func (t Type) methodName(parameter list) { } ``` 上面的代码片段创建了一个接收器类型为 `Type` ...

- Go 系列教程 —— 16. 结构体
  - 欢迎来到 [Golang 系列教程](/subject/2)的第 16 个教程。 ### 什么是结构体？ 结构体是用户定义的类型，表示若干个字段（Field）的集合。有时应该把数据整合在一起，而不是让这些数据没有联系。这种情况下可以使用结构体。 例如，一个职员有 `firstName`、`lastName` 和 `age` 三个属性，而把这些属性组合在一个结构体 `employee` 中就很合理。 ### 结构体的声明 ```go type Employee ...

- Go 系列教程 —— 15. 指针
  - 欢迎来到 [Golang 系列教程](/subject/2)的第 15 个教程。 ### 什么是指针？ 指针是一种存储变量内存地址（Memory Address）的变量。 ![指针示意图](https://raw.githubusercontent.com/studygolang/gctt-images/master/golang-series/pointer-explained.png "指针示意图") 如上图所示，变量 `b` 的值为 `156`，而 `b` 的内...

- Go 系列教程 —— 14. 字符串
  - 欢迎阅读 [Golang 系列教程](/subject/2)第 14 部分。 由于和其他语言相比，字符串在 Go 语言中有着自己特殊的实现，因此在这里需要被特别提出来。 ## 什么是字符串？ Go 语言中的字符串是一个字节切片。把内容放在双引号""之间，我们可以创建一个字符串。让我们来看一个创建并打印字符串的简单示例。 ```go package main import ( "fmt" ) func main() { name := "...

- Go 系列教程 —— 13. Maps
  - 欢迎来到 [Golang 系列教程](/subject/2)的第 13 个教程。 ## 什么是 map ？ map 是在 Go 中将值（value）与键（key）关联的内置类型。通过相应的键可以获取到值。 ## 如何创建 map ？ 通过向 `make` 函数传入键和值的类型，可以创建 map。`make(map[type of key]type of value)` 是创建 map 的语法。 ```go personSalary := make(map[str...

- Go 系列教程 —— 12. 可变参数函数
  - 欢迎来到 [Golang 系列教程](/subject/2)第 12 章。 ## 什么是可变参数函数 可变参数函数是一种参数个数可变的函数。 ## 语法 如果函数最后一个参数被记作 `...T ` ，这时函数可以接受任意个 `T` 类型参数作为最后一个参数。 请注意只有函数的最后一个参数才允许是可变的。 ## 通过一些例子理解可变参数函数如何工作 你是否曾经想过 append 函数是如何将任意个参数值加入到切片中的。这样 append 函数可以接受不同数...
 
- Go 系列教程 —— 11. 数组和切片
  - 欢迎来到 [Golang 系列教程](/subject/2)的第 11 章。在本章教程中，我们将讨论 Go 语言中的数组和切片。 ## 数组 数组是同一类型元素的集合。例如，整数集合 5,8,9,79,76 形成一个数组。Go 语言中不允许混合不同类型的元素，例如包含字符串和整数的数组。（译者注：当然，如果是 interface{} 类型数组，可以包含任意类型） ### 数组的声明 一个数组的表示形式为 `n[T]`。`n` 表示数组中元素的数量，`T` 代表每个元素的类...

- Go 系列教程 —— 10. switch 语句
  - 这是 [Golang 系列教程](https://golangbot.com/learn-golang-series/)中的第 10 篇。 switch 是一个条件语句，用于将表达式的值与可能匹配的选项列表进行比较，并根据匹配情况执行相应的代码块。它可以被认为是替代多个 `if else` 子句的常用方式。 看代码比文字更容易理解。让我们从一个简单的例子开始，它将把一个手指的编号作为输入，然后输出该手指对应的名字。比如 0 是拇指，1 是食指等等。 ```go package...

- Go 系列教程 —— 9. 循环
  - 这是 Go 语言系列教程的第 9 部分。 循环语句是用来重复执行某一段代码。 `for` 是 Go 语言唯一的循环语句。Go 语言中并没有其他语言比如 C 语言中的 `while` 和 `do while` 循环。 ## for 循环语法 ```go for initialisation; condition; post { } ``` 初始化语句只执行一次。循环初始化后，将检查循环条件。如果条件的计算结果为 `true` ，则 `{}` 内的循环体将在 ...

- Go 系列教程 —— 8. if-else 语句
 - 这是我们 [Golang 系列教程](https://golangbot.com/learn-golang-series/)的第 8 篇。 if 是条件语句。if 语句的语法是 ```go if condition { } ``` 如果 `condition` 为真，则执行 `{` 和 `}` 之间的代码。 不同于其他语言，例如 C 语言，Go 语言里的 `{ }` 是必要的，即使在 `{ }` 之间只有一条语句。 if 语句还有可选的 `else i...

- Go 系列教程 —— 7. 包
  - 这是 Golang 系列教程的第 7 个教程。 ### 什么是包，为什么使用包？ 到目前为止，我们看到的 Go 程序都只有一个文件，文件里包含一个 main [函数](https://studygolang.com/articles/11892)和几个其他的函数。在实际中，这种把所有源代码编写在一个文件的方法并不好用。以这种方式编写，代码的重用和维护都会很困难。而包（Package）解决了这样的问题。 **包用于组织 Go 源代码，提供了更好的可重用性与可读性**。由于包提供了...

- Go 系列教程 —— 6. 函数（Function）
  - 这是我们 [Golang 系列教程](/subject/2)第 6 章，学习 Golang 函数的相关知识。 ## 函数是什么？ 函数是一块执行特定任务的代码。一个函数是在输入源基础上，通过执行一系列的算法，生成预期的输出。 ## 函数的声明 在 Go 语言中，函数声明通用语法如下： ```go func functionname(parametername type) returntype { // 函数体（具体实现的功能） } ``` ...

- Go 系列教程 —— 5. 常量
  - 这是我们 Golang 系列教程的第 5 篇。 ### 定义 在 Go 语言中，术语"常量"用于表示固定的值。比如 `5` 、`-89`、 `I love Go`、`67.89` 等等。 看看下面的代码: ```go var a int = 50 var b string = "I love Go" ``` **在上面的代码中，变量 `a` 和 `b` 分别被赋值为常量 `50` 和 `I love GO`**。关键字 `const` 被用于表示常量，比如...

- Go 系列教程 —— 4. 类型
  - 这是我们 Golang 系列教程的第 4 个教程。 请阅读本系列的 [Golang 教程第 3 部分：变量](/articles/11756) 来学习变量的知识。 下面是 Go 支持的基本类型： - bool - 数字类型 - int8, int16, int32, int64, int - uint8, uint16, uint32, uint64, uint - float32, float64 - complex64, complex128 - by...


- Go 系列教程 —— 3. 变量
  - 这是我们 Golang 系列教程的第 3 个教程，探讨 Golang 里的变量（Variables）。 你可以阅读 Golang 系列 [**教程第 2 部分：Hello World**](/articles/11755)，学习如何配置 Golang，并运行 Hello World 程序。 ### 变量是什么 变量指定了某存储单元（Memory Location）的名称，该存储单元会存储特定类型的值。在 Go 中，有多种语法用于声明变量。 ### 声明单个变量 **var ...

- Go 系列教程 —— 2. Hello World
  - 这是 Golang 系列教程的第 2 个教程。如果想要了解什么是 Golang，以及如何安装 Golang，请阅读 [Golang 教程第 1 部分：介绍与安装](/articles/11706)。 学习一种编程语言的最好方法就是去动手实践，编写代码。让我们开始编写第一个 Go 程序吧。 我个人推荐使用安装了 [Go 扩展](https://marketplace.visualstudio.com/items?itemName=lukehoban.Go)的 [Visual Studi...

- Go 系列教程 —— 1. 介绍与安装
  - 这是我们 Golang 系列教程的第一个教程。 ## Golang 是什么 Go 亦称为 Golang (译注：按照 Rob Pike 说法，语言叫做 Go，Golang 只是官方网站的网址)，是由谷歌开发的一个开源的编译型的静态语言。 Golang 的主要关注点是使得高可用性和可扩展性的 Web 应用的开发变得简便容易。（译注：Go 的定位是系统编程语言，只是对 Web 开发支持较好） ## 为何选择 Golang 既然有很多其他编程语言可以做同样的工作，如 Python...

