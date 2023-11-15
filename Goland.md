# 1 前言

梦想golang工程师...

* 前置要求：

1. 后端编程语言 ok

2. 网络编程+并发思想 ok

3. 计算机基本体系结构 ok

4. Linux基本使用 ok

冲！

# 2 开发环境和IDE

* 源码包的下载

<https://golang.google.cn/dl/>

<https://studygolang.com/dl>

***

* 解压

```shell
sudo tar -zxvf go1.23.3 linux-arm64.tar.gz -C /usr/local/
# 查看go源码包(go编译器)
cd /usr/local/go
```

`/usr/local/go`下有一些目录：

* `src` 

go的源码

* `bin`

默认有两个指令，`go`和`gofmt`，相当于`gcc`等命令

所以要将该路径置于环境变量中：

```shell
vim ~/.bashrc
G # 到文件末尾
# 添加以下内容

# 源码包所在路径
export GOROOT=/usr/local/go
# 当前用户ty，写go语言用的工作路径
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
# 保存
:wq
# 加载
source ~/.bashrc
# 测试
go version # 如果成功输出版本号，则配置成功
go --help
```

IDE 推荐:

Goland vscode Vim

# 3 当下golang分析

## 3.1 优势

* 极其**简单的部署**方式
  
  * 可以直接编译成机器码
  
  * 不依赖其它库
  
  * 直接运行即可部署

* **静态类型**语言
  
  * 编译时检查出来隐藏的大多数问题

* 语言层面的**并发**
  
  * 天生的基因支持(区分天生和美容)
  
  * 充分利用多核

* 强大的标准库
  
  * runtime系统调度机制
  
  * 高效的GC垃圾回收
  
  * 丰富的标准库

* 简单易学
  
  * 25个关键字
  
  * C系语言，内嵌C支持
  
  * 面向对象
  
  * 跨平台

* 大厂领军

*** 

用途：

1. 云计算基础设施领域

2. 基础后端软件

3. 微服务

4. 互联网基础设施

成就：

docker kubernetes

## 3.2 不足

1. 包管理，大部分包都在github上

2. 无泛化类型，计划2.0加上

3. 无异常处理，通过`Error`来处理(类似C)

4. 对C的降级处理，并非无缝，没有C降级到asm(汇编)那么完美

# 4 go语法新奇

## 4.1 从一个main函数初始go

```go
package main // 程序的包名
/*
import "fmt"
import "time"
*/
import {
    "fmt"
    "time"
}
// go强制要求左花括号必须在函数名在同一行
func main() {
    // 没有分号，有无分号无所谓，建议不加
    fmt.Println("Hello, go!") 
}
```

```shell
go run hello.go # run 编译+运行
```

* 总结
  
  * 导包的两种方式
  
  * 左花括号不能占一行
  
  * 不加分号

## 4.2 常见的四种变量声明方式及多变量声明方式

```go
package main

import "fmt"

// 声明全局变量，方法1，2，3可以
var ga int
// 方法4 不可以

func main() {
    // 方法1 ：声明一个变量,默认值为0
    var a int

    // 方法2：
    var b int = 100

    // 方法3：初始化时省略类型说明，类似auto
    var c = 100
    // 打印变量类型
    fmt.Printf("type of c = %T\n", c) // 格式化输出 %T

    // 方法4(常用): 省略var关键字，直接自动匹配
    d := 100

    // 声明多个变量
    var x, y int = 100, 200
    var e, e = 100, "av"
    var {
        v1 int = 100
        v2 bool = true
    }
}
```

## 4.3 const和itoa注意事项

```go
package main

import "fmt"

// const 定义枚举类型
const (
    // 可以在 const() 添加一个关键字 iota
    // 第一行的iota默认是0，以下每行累加1
    /*
    red = 0
    green = 1
    blue = 2
    */ // 依次赋值比较麻烦
    red = iota
    green
    blue
    // 右值可以是含iota的公式，下面的值也会继承这个式子
)

// 例如
const (
    a = iota * 10 + 6 // 0 * 10 + 6
    b  // 1 * 10 + 6
    c  // 2 * 10 + 6
)

const (
    a, b = iota + 1, iota + 2   // iota = 0
    c, d                        //  iota = 1
    e, f                        // iota = 2
    g, h = iota * 2, iota * 3   // iota = 3
    i, k
)
// iota 只能用在const的括号中

func main() {
    // 常量 只读
    const len int = 10
}
```
## 4.4 golang中函数多返回值

* 定义函数

```go
func <fn-name>(args arg_t, ...) <return-val-type> {
    // fn-body
}
```

```go
// 普通函数
func f1(a string, b int) int {
    return 0
}

// 返回多个返回值，匿名
func f2(a string, b int) (int, int) {
    return 0, 1
}

// 返回多个返回值，有形参名称
// r1,r2是属于f3的两个形参，初始化默认为0
// r1,r2的作用域就是f3的函数体
func f3(a string, b int) (r1 int, r2 int) {
    // 在未赋值之前 r1 = r2 = 0
    // 给有名称的返回值变量赋值
    r1 = 100
    r2 = 200
    return
}

// 多个返回值同种类型，形参同理
func f4(a string, b int) (r1, r2 int) {
    return 100, 200 // 也ok
}
```

## 4.5 import导包路径问题和init方法调用流程

`init()`方法可以初始化当前包的一些资源，优先于`main()`的调用:

![](image/process.png)

一般一个包对应一个文件夹：

```
|
|---lib1|
|       |---lib1.go
|
|---lib2|
|       |---lib2.go
|
|---main.go
```

一般包名于文件名一致

```go
// lib1.go
package lib1
import "fmt"
func init() {
    fmt.Println("lib1 init()")
}

// 函数名首字母大写，表示对外开放
func Test() {
    
}
// lib2.go 同1
// main.go
package main
import (
    "lib1" // 指定从GOPATH开始到当前包的路径
    "lib2"
)
func main() {

}
```

## 4.6 import匿名及别名导包方式

在Go中，如果导入了一个包，那么要至少使用一个这个包中的API

考虑一种情况，只想使用这个包的`init()`方法，而不想使用其它方法

此时可在包名前使用`_ `表示匿名别名

同时这个别名也可以自己定义

```go
import (
    _ "path/lib1"  // 无法使用当前包的方法，但可以使用init方法
    
    ty "path/lib2" // 此时使用 ty.
    
    . "path/lib3" // 此时使用不需要 . 相当于导出namespace
    // 同样少用这种
)
func main() {
    ty.Test()
    TestLib3()
}
```

## 4.7 指针速通

熟悉C/C++直接跳过

go中指针与C/C++还是有一些区别的：

* 结构类型指针访问成员，在go中一致用`.`而不是`->`

* 指针算术运算：go不允许指针算术运算，go中指针主要用于引用和传址

* 相对安全

## 4.8 defer语句调用顺序

```go
func f1() int{
    fmt.Println("f1")
    return 0
}
func f2() int{
    fmt.Println("f2")
    return 0
}
func test() {
    defer f1()
    return f2()
    // 先return 后defer
    // 当前函数声明周期全部结束后defer才执行
    // 所以要在return之后
}

func main() {
    // defer expr -- 当前作用域结束或当前流程之后自动调用
    // 类似 C++ 的 析构函数
    defer fmt.Println("done")
    // defer 可以有多个 执行按照压栈顺序 -- 先定义的defer后执行
    f1()
    f2()
}
```

## 4.9 数组/slice(动态数组)

* **静态数组**

```go
// 静态数组传参
// 值拷贝并且只能传10大小的数组
func printArray(arr [10]int) { 
    
}

func main() {
    // 固定长度的数组，默认值都为0，相当于std::array
    var arr [10] int
    arr1 := [10]int{1, 2, 3, 4}
    
    // 遍历
    // 1 索引for (注意：go中只有后置++)
    for i := 0; i < len(arr); i++ {

    }
    // 2 范围for 关键字range  返回索引和值
    for index, value := range arr1 {

    }

    // 查看数组的数据类型
    fmt.Printf("arr type is %T\n", arr) // [10]int
}
```

* **动态数组/切片**

```go
// 引用传递
func showSlice(slc []int) {
    // _ 表示匿名变量
    for(_, val := range slc) {
        fmt.Println(val)
    }
}

func main() {
    // 声明
    slc := []int {1, 2, 3, 4}  // slice --- []int
}

```

### 4.9.1 slice--声明方式

声明一个切片有以下4种方式
```go
func main() {
    // 1 
    slice1 := []int {1, 2, 3}

    // 2 声明一个切片，但是没有给它分配空间
    var slice2 []int
    // 开辟空间,默认值为0
    slice2 = make([]int, 3)

    // 3 (2)可以合在一起
    var slice3 = make([]int, 3)

    // 4 通过 := 推导出slice4是一个切片
    slice4 := make([]int, 3)

    // %v --- 表示打印详细信息
    fmt.Printf("len = %d, slice = %v\n", len(slice1), slien1)
    // len = 3, slice = [1, 2, 3]

    // 判断一个slice(的内存空间)是否为空
    if slice1 == nil {
        fmt.Println("This is a empty slice.")
    } else { // 要在一起，否则有语法错误
        fmt.Println("It is not empty.")
    }
}
```

### 4.9.2 slice --- 使用方式(追加和截取)

使用`make()`返回一个切片时可以指定该切片的容量(capacity), 

这里要注意区分大小(size)和容量(capacity)

```go
func main() {
    // 定义一个长度为3，容量为5的切片
    nums := make([]int, 3, 5)
    // 求长度 len()
    // 求容量 cap()

    // 向 nums 切片追加 1
    nums = append(nums, 1)

    // 若切片容量已满，调用append()会将cap扩大到原来的2倍

    nums1 := nums[0:2] // 左闭右开，但这样只是复制了指针

    // copy() 可以将底层的slice一起进行拷贝 深拷贝
    // copy(dest, src)
}
```

## 4.10 map(hash)

### 4.10.1 声明方式

定义一个`map`

```go
// ---------第一种声明方式-----------------
// 声明mp是一种map类型，key是string，value是string
var mp1 map[string]string // mp1 == nil

mp1 = make(map[string]string, 10)
mp1["one"] = "cpp"
mp1["two"] = "python" // 自动扩容
fmt.Println(mp1)  // 打印出来无序，是哈希映射

// ---------第二种声明方式-----------------
mp2 := make(map[int]string)
mp2[1] = "cpp"
mp2[4] = "ccc"

// ---------第三种声明方式-----------------
mp3 := map[int]string {
    1 : "php",
    2 : "c++", // 最后一个也要有逗号
}
```

### 4.10.2 使用方式

```go
mp := make(map[string]int)
// 添加 operator[]

// 遍历
for key, value := range mp {

}

// 删除 , 参1 --- map对象; 参2 --- key
delete(mp, 2)

// 修改 operator[]

// 引用传递
func test(mp map[string]int) {

}
```

## 4.11 struct --- obj


关键字`type`，声明一种(新的)类型

```go
type myint int // 类型重定义 --- typedef

// 定义结构
type book struct {
    title string
    author string
}

func test() {
    var book1 Book
    // 访问成员
    book1.title = "Golang"
    book2.title = "ttt"
}

// 结构传递时属于 值传递
func test1(book Book) {

}

// 可以传指针
func test2(book *Book) {

}
```

## 4.12 objection

### 4.12.1 封装

注意go是根据包来进行封装的，比如：

有一个类，它的字段名小写，但当前包依旧能访问其私有字段

区别C++按照`class`进行封装

```go
// 如果类名首字母大写，表示其它包也能够访问
type Hero struct {
    // 如果类的属性名首字母大写，表示public
    //                   小写      private
    name string
    ad int
    level int
}

// 结构绑定方法
/*
func (this Hero) getName() {
    // 当前this是调用该方法对象的一个拷贝
    fmt.Println(this.name)
}
*/
// 名字不一定是 this
func (this *Hero) getName() {
    // 当前this是调用该方法对象的地址
    fmt.Println(this.name)
}

func test() {
    // 创建一个对象
    hero := Hero(name:"aaa", ad:100, level:1)

    // 调用成员函数
    hero.getName()
}

```

### 4.12.2 继承

```go
type Human struct {
    name string
    sex bool
}

func (this *Human)Eat() {

}

type SuperMan struct {
    Human // 仅有一个类型名，表示当前类继承了这个类
    // 直接包含基类所有属性，不分啥public，protected...
    level int
}

// 重定义/重写 父类方法 Eat()
func (this *SuperMan) Eat() {

}

// 子类新方法
func (this *SuperMan) Fly() {

}

func main() {
    h := Human{"zhang3", true}
    h.Eat()
    // 定义子类对象
    s := SuperMan{Huamn{"li3", false}, 88}
    var s SuperMan
    s.name = "li3"
    s.sex = false
    s.level = 88
    s.Eat()
    s.Fly()
}
```

### 4.12.3 多态

**interface**

// interface 本质是一个**指针**
// 保存了当前interface执行的具体类型和绑定的函数列表

```go
type animal interface {
    Sleep()
    GetColor()
    GetType()
}

// 具体的类
type Cat struct {
    // 不必指定继承 animal1,直接重写方法就好
    color string
}

func (this *Cat) Sleep() {

}
// ..其余同理

// 重写3个方法，要重写接口的全部方法

func main() {
    var animal animal_interface
    animal_interface = &Cat{"green"}
    animal_interface.Sleep()
}
```

## 4.13 interface空接口万能类型和断言

通用万能类型 `interface{}` --- 空接口

基本类型如int、float、struct等都实现了`interface{}`

```go
func fn(arg interface{}) {
    fmt.Println(arg)
}
```

使用上类似C中的`void*`,不过省去了类型转换这一步

* 给`interface{}`提供类型断言机制

```go
val, ok : arg.(string)
if !ok {
    // not a string
} else {
    fmt.pritf("%s", val)
}
```

## 4.14 变量的内置pair结构详细说明

每个变量内部保存一个对组`pair`保存该变量的类型和值

type(static type || concrete type) + value

static type : int,string...(基本类型)

concrete type : interface **指向**的具体数据类型，具体类型

在该变量传递过程中，这个对组保存的类型和值不会发生改变

```go
import "fmt"

type Reader interface {
    ReadBook()
}

type Writer interface {
    WriteBook()
}

// 具体类型
type Book struct {
}

func (self *Book) ReadBook() {
    fmt.Println("read a book")
}

func (self *Book) WriteBook() {
    fmt.Println("write a book")
}

func main() {
    // b:pair<type:Book, value:&Book{}>
    b := &Book{}
    
    // r:pair<type:,value:> 均为空
    var r Reader
    // r:pair<type:Book, value:&Book{}>
    r = b
    r.ReadBook()

    var w Writer
    // w:pair<type:Book, value:&Book{}>
    w = r.(Writer) // 此处断言成功：因为w和r具体的type是一致的
}
```


## 4.15 反射reflect

reflect：已知当前变量，得到当前变量的`type`和`value`

reflect包：

```go
func ValueOf(i interface{}) Value {...}
// ValueOf 用来获取输入参数接口中数据的值，如果接口为空则返回0

func TypeOf(i interface{}) Type {...}
// TypeOf 用来动态获取输入参数接口中值的类型，如果接口为空返回nil
```

### 4.15.1 反射一个类型

```go
package main
import {
    "fmt"
    "reflect"
}

func reflectFloat(arg interface{}) {
    fmt.Println("type:", reflect.TypeOf(arg))
    fmt.Println("value:", reflect.ValueOf(arg))
}

type User struct {
    id int
    name string
    age int
}

func (self *User) Call() {

}

func main() {
    var num float64 = 1.2345
    reflectFloat(num)
    // float 1.2345

    user := User{1, "ace", 18}
    doFieldAndMethod(user)
}

func doFieldAndMethod(input interface{}) {
    inputType := reflect.TypeOf(input)
    inputValue := reflect.ValueOf(input)
    fmt.Println(inputType.Name()) // user
    fmt.Println(inputValue) // ...

    // 通过type获得里面的字段
    // 1.通过inputType得到其中的字段个数NumField
    // 2.得到每个field，数据类型
    // 3.通过field有一个Interface()方法得到对应value
    for i := 0; i < inputType.NumField(); i++ {
        field := imputType.Field(i)
        value := imputValue(i).Interface()

        fmt.Printf("%s : %v = %v\n", field.Name, field.Type, value)
        /*
        id : int = 1
        name : string = ace
        age : int = 18
        */
    }

    // 通过type获得里面的字段
    for i := 0; i < inputType.NumMethod(); i++ {
        m := inputType.Method(i)
        fmt.Printf("%s : %v\n", m.Name, m.Type)
        // Call : func(main.User)
    }
}
```

### 4.15.2 反射解析结构体标签Tag

golang允许在结构体字段后加标签

标签为键值对,用一对反单引号包裹，可以有多个，中间用空间隔开

标签的键应为**已知信息**，值为具体说明

主要作为导包时的一些说明

```go
type resume struct {
    Name string `info:"name" doc:"我的名字"`
    Sex string `info:"sex"`
}

func findTag(sru interface{}) {
    // Elem()表示当前结构体的全部元素
    t := reflect.Typeof(sru).Elem() 
    
    for i := 0; i < t.NumField(); i++ {
        taginfo := t.Field(i).Tag.Get("info")
        tagdoc := t.Field(i).Tag.Get("doc")
        fmt.Println("info:" taginfo, "doc:" tagdoc)
    }
}
```

应用：解析Json文件,orm映射关系

### 4.15.3 结构体标签在Json的应用

```go
import (
    "encoding/json"
    "fmt"
)

// 用于 Json文件
// 标签Tag的key用来表示json，不可修改
// 标签Tag的value用来作为json文件中的键，可修改
type Movie struct {
    TiTle string    `json:"title"`
    Year int        `json:"year"`
    Price int       `json:price`
}

func main() {
    movie := Movie{"喜剧之王", 2000, 8};

    // 编码: struct --> json
    jsonStr, err := json.Marshal(movie);
    if err != nil {
        // error
        return;
    }

    fmt.Println(jsonStr)

    // 解码：json --> struct
    my_movei := Movie{}
    err = json.Unmarshal(jsonStr, &my_movie)
    if err != nil {
        return
    }
    fmt.Println(my_movie)
}
```

一般的话库里会提供相关API解析结构体标签


# 5 golang高阶

## 5.1 goroutine基本模型和调度设计策略

### 5.1.1 协程初识 

早期**单进程**操作系统 --> 线程顺序执行

* 单进程时代的两个问题：
  
  1. 单一的执行流程，计算机只能一个任务一个任务处理
  
  2.  进程阻塞所带来的CPU浪费时间


* 能不能宏观的执行多个任务？ --> 多线程/多进程操作系统

CPU调度器(轮询调度) ---- 时间片(一个进程/线程一次执行的最大时间)

即并发执行

显然：多线程/多进程解决了阻塞问题

但是，进程或线程的切换是有成本的 ---> cpu时间浪费

因此：进程/线程的**数量越多**，切换成本就**越大**，也就越**浪费**

提高**CPU的利用率**，才是当今软件层面上系统优化和架构要做的事

多线程 随着**同步竞争**(如锁、竞争资源冲突) 使开发设计变得越来越复杂

此外，线程和进程还是比较浪费内存的...

总结来看，多线程多进程存在以下问题

![](image/concurrentShortage.png)

提高CPU利用率，同时解决高消耗调度、高内存就是下一代并发语言或设计的重点

*** 

* 协程(co-routine)

一个*线程*会被操作系统分为**用户态**和**内核态**

![](image/threadStructure.png)

* 内核态：表示操作系统底层，包括
  * 线程开辟
  * 物理内存资源分配
  * 磁盘资源分配

* 用户态：上层开发写业务逻辑的，调接口写应用程序

考虑将用户态的内核态一分为二：对应用户线程和内核线程，并将它们绑定起来：

![](image/twoThread.png)

其中，内核线程**单独地**整理硬件层面的东西

而用户线程来保证业务层面上的并发效果

绑定完之后，CPU的视野就只关心内核线程，也就是说：操作系统是不用改变的

然后给它们改个名称：

内核线程依然叫做线程(thread)

用户线程起个别名叫**协程(co-routine)**

类比线程调度，那么此时一个(内核)线程可不可以绑定一个**协程调度器**来绑定多个协程

此时CPU依旧无感(它的视野里中只有内核线程)，在上层开了多个用户线程(协程)

每个协程可以挂载一个任务，这样：

用户态依然能保证并发效果，而CPU本身不需要切换

这样就解决了高消耗CPU调度瓶颈，这种方案显然是可以的

称为`N:1`协程对应关系

这种方式也是有一定弊端的：

如果一个协程阻塞了，那么可能耽误下一个协程调度：

![](image/N_1.png "N:1-关系")

***

那么改变为`1:1`关系呢？虽然解决了阻塞，但又回到线程级别

即此时协程的创建、删除和切换的**代价**都是由CPU完成的

***

* `M:N`关系

多个线程通过协程调度器管理读个协程：

![](image/M_N.png "M:N关系")

`M:N`关系将全部的瓶颈方法集中在优化**协程调度器**上

协程调度器越好，CPU利用率就越高

不同的编程语言，想要支持协程，就要实现自己的协程调度器

谁的协程调度器做得好，谁对协程的支持就更加好一些

因为底层CPU的调度是我们优化不了的

### 5.1.2 GoLang协程---Goroutine

* 对协程的处理

对协程调度器处理之前，go先对协程做了一些处理

首先是命名，在go中，协程(co-routine)被叫做`Goroutine`

另外是内存 ：几KB (相比于线程几MB)，因此可以大量存在

还有：灵活调度(调度器方面) ---> 可常切换

***

* **Golang对早期调度器的处理**

![](image/earlyScheduler.png "Go早期调度器")

上图中，`M`表示线程，`G`表示协程

存在若干个线程

有一个全局的协程队列，每创建一个协程，就把它放到该队中

当`M0`想要调用一个`Gotoutine`时，首先尝试获取全局队列的锁的所有权

获取锁之后，执行其中一个`Goroutine`，首先会将要执行的`Goroutine`出队，

剩余没有被执行`Goroutine`的前移，移动至队头，以便下次的`M`获取

执行完一个`Goroutine`，将这个`Goroutine`放回队列

* 老调度器的弊端

老调度器的实现时非常简单的，它里面有很多弊端：

1. 创建、销毁、调度`G`都需要`M`获取锁，这样就形成了**激烈的锁竞争**

2. `M`转移`G`会造成**延迟和额外的系统负载**

3. 系统调用(CPU在`M`之间的切换)导致频繁的线程阻塞和取消阻塞操作增加了系统开销

解释2：

假设`G`中又创建了一个新的协程`G_`，为了保证并发，`G_`也应被执行

但是`G_`可能被其它`M`获取执行，这就导致了很差的**局部性**

局部性：`G`在`M`中执行，那么就希望`G`生成的`G_`也在`M`中执行


* **GMP模型**

G --- goroutine 协程

P --- processor 处理器(管理协程)

M --- thread 线程

processer 包含了每个协程的资源，如果要运行一个协程

那么首先要先获取`P`

![](image/GMP.png "GMP Model")

操作系统调度器：用来调度CPU

用户态：内核线程之上

`P`中包含了每个协程的资源，包括堆栈数据等，`P`最多`GOMAXPROCS`(一个宏)个

P的本地队列(local p queue)，每个`P`都会有

除了本地队列以外，还有一个全局队列，存放等待运行的`G`(实际上本地队列也是)

新创建的`G`优先放在本地队列中，若本地队列全部满了，那么放到全局队列中

一个`P`同一时刻只能运行一个`Goroutine`

当前程序最大并行数量即`GOMAXPROCS`

### 5.1.3 调度器的设计策略

* **复用线程**

**work steading机制** 

偷取 --- 当某个`P`的本地队列为空且不工作时，它可以从别的`P`的队列中拿来等待的`G`

**hand off机制** 

分离 --- 某个正在执行的`G`阻塞后，让其它协程一直等待是没有意义的

此时就会创建或唤醒一个物理线程`thread`然后阻塞的`G`所在的`P`会与当前线程分离

与新的线程结合，这样就不因一个协程的阻塞而影响同队中后续协程的执行

此时这个旧的线程与阻塞的协程结合(由于协程阻塞，它并不工作即睡眠)

若这个阻塞的协程执行完毕，还要继续执行的话就会加入其它队列中

如果不执行则该线程就会睡眠或销毁

* 利用并行

`GOMAXPROCS`(CPU核心数/2)限定`P`的个数

* **抢占**

co-routine:原先一个协程绑定了CPU，如果它不主动释放，那么后续协程只能永远等待

goroutine:后续的协程最多等待先前的10ms，若超时，直接抢占CPU

保证平等并发

* 全局G队列

基于Work stealing做的补充

如果某个`P`的本地队列为空且对应线程空闲，那么这个`P`优先偷其它`P`的本地队列中的协程`G`

如果其它`P`的本地队列也空了，那么它会从全局队列中偷(需要加锁和解锁)

## 5.2 创建goroutine

通过关键字`go`创建一个`goroutine`

`go`后面紧接一个函数调用

```go
func task() {

}

func main() {
    // 创建一个go程，执行task()流程
    go task()
    // 主线程继续继续

    // 为了防止主线程退出，这里做一个死循环
    i := 0
    for {
        i++
        ...
    }
    // 主线程退出，子协程也会销毁


    // 匿名函数作为任务函数
    go func() { // 匿名函数 形参为空 返回值空
        defer fmt.Println("a. defer")

        // 在定义一个匿名函数并调用
        func() {
            defer fmt.Println("b.defer")
            fmt.Println("b")
        }()

        fmt.Println("a")
    }()
}
```

go程的退出需要任务函数退出，在任务函数中的子函数`return`不会退出这个go程

如果要在任务函数中调用的函数中退出当前go程，则需要调用一个函数`runtime.Goexit()`

获取子go程任务函数的返回值，由于不同go程间属于异步操作

因此需要借助`channel`管道机制

## 5.3 channel的基本定义和使用

两个`goroutine`间相互通行需要借助`channel`(管道)，类似`ipc`

* 定义`channel`
```go
make(chan Type) // 无缓冲，等价于 make(chan Type, 0)
make(chan Type, capacity)
```

* 相关操作

```go
// 发送value到channel
channel <- value
// 接收数据并将其丢弃
<- channel
// 接收数据并赋值给x
x := <- channel
// 功能同上，同时检测通道是否已关闭或空
x, ok := <- channel
```

```go
func main() {
    // 定义一个channel
    c := make(chan int)
    go func() {
        defer fmt.Println("goroutine finished")
        fmt.Println("goroutine is working")
        c <- 666 // 将666写入c   ***
    }()

    num := <- c // 从c中接收数据， ***
    fmt.Println(num)
    fmt.Println("main goroutine finished")
}
```

多运行几次会发现：基本上都会时`main goroutine`最后结束

`channel`本身具有同步两个不同go程的能力:

比如在上述实例中：

若main(go程)先执行到`***`，那么通道就会阻塞该go程，直到子go程执行完`***`

若子go程先执行到`***`，通道同样阻塞该go程，直到main-go程执行到`***`

## 5.4 channel有无缓冲同步问题

### 5.4.1 无缓冲channel

通道中不允许存放数据，即上节中的示例


### 5.4.2 有缓冲channel

通道中可以存放数据

若通道不空，则取数据不会被阻塞；若通道不满，则放数据不会被阻塞

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 有缓冲的通道
    c := make(chan int, 3)
    // 显示元素个数len和容量cap
    fmt.Println(len(c), " ", cap(c))

    go func() {
        defer fmt.Println("func finish")

        for i := 0; i < 3; i++ {
            c <- i
            fmt.Println("func work, i = ", i)
        }
    }()
    
    time.Sleep(2 * time.second)

    for i := 0; i < 3; i++ {
        num := <-c // 此处的箭头和变量名之间不能加空格
        fmt.Println(num)
    }
    fmt.Println("func finish")
}
```

## 5.5 channel的关闭特点

关闭通道 ： `close(c) // c := make(chan Tp, cap)`

`if`语句扩展：

```go
if expr; cond {

}

// 其中 expr为表达式，cond为判定条件
```

```go
func main {
    c := make(chan int)

    go func() {
        for i := 0; i < 5; i++ {
            c <- i
        }
        // 关闭通道
        close(c)
    }()
    for {
        if data, ok := <-c; ok {
            fmt.Println(data)
        } else {
            break;
        }
    }
}
```

上述示例中，子go程写完数据后必须要关闭通道：

否则主go程会一直尝试读空的通道

* 注意：

1. `channel`不像文件一样需要经常关闭，只有确认没有发送数据或要显示结束循环才关闭

2. 关闭`channel`后，无法向`channel`在发送数据(引发panic错误，返回0)

3. 关闭`channel`后，可以继续从中读数据

4. 对于`nil channel`，无论收发都会被阻塞

## 5.5 channel和range

上节示例中主go程中读数据是可以通过`range`关键字简化的：

```go
for data := range c {
    fmt.Println(data)
}
// 效果相同
```

## 5.6 channel和select

单流程下一个go只能监控一个channel，

但是通过关键字`select`可以完成监控**多个**`channel`的状态

就类似于IO多路复用中的`select()`和`epoll`

如果有个`channel`可读或可写就会返回

```go
select {
    case <-chan1:
        // 如果chan1成功读到数据，则进行该case处理语句
    case chan2 <- 1:
        // 如果成功向chan2写入数据，则进行该case处理语句
    default:
        // 如果上面都没有成功，则进入default处理流程
}
```

一般来说外层会嵌套一层`for`

示例：

```go
import "fmt"

func fib(c, quit chan int) {
    x, y := 1, 1;
    for {
        select {
            case c <- x:
                t := x
                x = y
                y = x + t
            case <- quit:
                fmt.Println("quit")
                return;
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)

    // sub go
    go func() {
        for i := 0; i < 6; i++ {
            // 读出写入通到中的数据
            fmt.Println(<-c)
        }
        // 发送停止信号
        quit <- 0
    }
    
    // main go
    fib(c, quit)
}

```

`select`具备多路`channel`的监控功能


# 6 go modules模块管理

`Go modules`是Go语言的依赖解决方案，正式于Go1.14,推荐使用

Go modules目前集中在Go的工具链中，只要安装了Go就可使用

Go modules的出现也解决了Go1.11前的几个常见争议问题：

1. Go语言长久依赖的依赖管理问题

2. "淘汰"现有的GOPATH使用模式

3. 统一社区中其它依赖管理工具，提供迁移功能

## 6.1 GOPATH工作模式的弊端

1. 没有版本控制概念 --- `go get`不能指定版本号

2. 无法同步一致第三方版本号

3. 无法指定当前项目引用的版本号 

## 6.2 GoModules模式基础环境说明

### 6.2.1 指令

* 帮助：查看所有`go mod`可用指令

`go mod help`


* 常用指令

```shell
go mod init  #生成go.mod文件

go mod download # 下载go.mod文件中指明的所有依赖

go mod tidy # 整理现有的依赖

go mod graph # 查看现有的依赖结构

go mod edit # 编辑go.mod文件

go mod verdor # 导出项目中所有的引来到vendor目录 

go mod verify # 校验一个模块是否被篡改过

go mod why # 查看为什么需要依赖某模块
```

### 6.2.2 相关环境变量

可以通过`go env`命令来查看


```bash
$ go env
```

![](image/goenv.png "go env")

**GO111MODULE**

这个环境变量是Go Modules的开关，其允许设置以下参数：

* auto：只要项目包含了go.mod文件就会启用Go modules

* on：启用Go modules，推荐设置

* off：禁用Go modules，不推荐设置

可以通过以下执行来设置：
```bash
$ go env -w GO111MODULE=on
```

**GOPROXY**

这个环境变量主要用于设置Go模块代理(Go module proxy)

其作用是用于使Go在后续拉取模块版本时直接通过镜像站点快速拉取

其默认值时：`https://proxy.golang.org,direct` 需要梯子

因此可以设置国内的代理：

* 阿里云 `https://mirrors.aliyun.com/goproxy/`

* 七牛云 `https://goproxy.cn,direct`

修改：
```bash
$ go env -w GOPROXY=https://goproxy.cn,direct
```

`direct` -- 指示符

当指定下载某第三方的包时，优先从代理中下载，

如果没有，再从源中下载


**GOSUMDB**

保证拉取的第三方库是完整的，没有被篡改的

如果设置了`GOPROXY`就不用在设置这个了.

**GONOPROXY/GONOSUMDB/GOPRIVATE**

一般设置`GOPRIVATE`即可

表示当前库是一个私用库，不需要从代理拉取和校验

## 6.3 GoModules初始化项目

通过`go mod init`来初始化项目

若成功，会生成一个`go.mod`文件

![](image/mod_init.png "init")

写一个测试文件 拉取依赖

![](image/go_get.png "go-get")


总结：

1. 开启Go Modules模块，设置`GO111MODULE=on`
2. 初始化项目 
   
   1. 任意文件夹创建一个项目(不要求在$GOPATH/src)
   
   2. 创建go.mod文件，同时起当前项目的模块名称
   
   3. 生成一个go.mod文件(模块名称+编译器版本号)
   
   4. 在该项目下写源代码，如果源代码中依赖某个库 `go get rul`
   
   5. go.mod文件会新增一行代码，含义：当前模块依赖的模块版本号
   
   6. 会生成一个go.sum文件：
    
    * 罗列当前项目直接或间接依赖的所有版本,保证今后项目依赖的版本不会被篡改
    
    * h1:hash 表示整体项目的zip文件打开之后的全部文件的校验和来生成的hash
    
    * 如果h1:hash不存在，可能表示依赖的库用不上
    
    * xxx/go.mod h1:hash --- go.mod文件做的hash

## 6.4 改变模块的依赖关系

```shell
go mod edit -repalce=<old-version>=<new-version>
```

# 7 案例-即时通信系统

架构图(简略)：

![](image/base.png)

server: onlineMap(用户信息) channel(广播)

user: 读写go程分离

## 7.1 构建基础Server

`~/go/src/golang-im-system`

### 7.1.1 gomod导入本地包

基础的Server用go实现相较于C是简单了很多的

这里遇到的主要的问题是go导包的问题

可能是因为GOPATH的工作方式逐渐被淘汰，因此ycm不支持提示

所以这里尝试使用gomod提高编码体验：

***

首先是`go mod init`初始化项目

然后在`main.go`中导入当前文件的**其他包**：(此处要**特别注意**)

* GOPATH

在这个模式下`import`语句指定的是文件路径，例如：

![](image/gopath_import.png)

不知道怎么的，我截图测试的时候这个煞笔gopath又有提示了它mua的

* GOMOD

在这个模式下指定的是包的嵌套关系(同一项目):

![](image/mod_iimport.png)

***

以上两种方式抛去智能提示不谈，程序都能跑起来！！！

另外要注意一点：

区别于C/C++当有变量定义却未使用时，ycm提示C/C++一个警告，而提示go一个错误

此外，当导了包没用的时候，也会提示一个错误，握草看着真难受一开始



## 7.2 用户上线功能

大致思路：

创建一个`User`类型的类来存储用户相关信息，`Server`由一个哈希表存储所有用户的指针

当有新连接时，向表中插入一个新的对组，同时向广播消息`Server.Message`中填写内容

`Message`应**不断**被监听，一旦有内容，就遍历哈希表发给每个用户

在由每个用户的方法发送给对应的客户短


## 7.3 用户消息广播机制

这个相同于`recv()`接收数据，太久没写都快忘了卧槽

## 7.4 用户业务层封装

这部分主要是重构代码，`Handler()`方法中很多操作可以是用户的方法

## 7.5 在线用户查询
## 7.6 修改用户名
## 7.7 超时强踢功能

`select`在执行过程中，若前面的`case`条件判断为真，

那么后续的`case`也会执行，不过不会进入`case`体中

```go
select {
    case <-actinve: // 1
        // 当前用户活跃，应重置定时器
        // 这里不做任何事情，为了激活select，更新下面的定时器
    case <-time.After(time.Second * 10) : // 2
        // 已经超时，释放资源并退出当前go程
}
```

解释：

进入`select`可以理解为**所有**`case`**同时**检测

如果`case1`触发，那么当前`select`就会返回，同时更新定时器`time.After`

如果指定时间内`case1`没有触发，`case2`就会执行，释放资源退出go程

## 7.8 私聊功能

1. 获取对方用户名

2. 根据用户名，得到对方User对象

3. 获得消息内容，通过对方的User对象将消息内容发送过去

## 7.9 客户端实现

### 7.9.1 建立连接

### 7.9.2 解析命令行

Go语言解析命令行 --- `flag`库

对命令行的解析应该在`main()`之前即`init()`中执行

```go
// 命令行中解析出来的ip
var serverIp string
//端口同理
var serverPort int
func init() {
    // flag.TypeVar()  Type为go中的变量类型
    flag.StringVar(&serverIp, "ip", "192.168.18.128", "Set server's ip.")
    // 参2--> 指定命令行参数 ./client -ip 192.168.18.128
    // 参3--> 默认值
    // 参4--> 说明--help、-h显示出的帮助
}


// init()中只是指明如何解析，要去解析还要再main中指定
func main() {
    // 命令行解析
    flag.Parse()
}
```

### 7.9.3 菜单显示

### 7.9.4 更新用户名

### 7.9.5 公聊模式


### 7.9.6 私聊模式

1. 查询当前都有哪些用户在线

2. 提示用户选择一个对象进入私聊

3. 发消息即可

待解决：有时候对方可能再选择时在线，发送消息时离线

## 7.10 后续完善

传输文件、视频、可视化UI...

# 8 Golang生态扩展介绍及未来成长方向

## 8.1 生态体系(知名)

1. **Web框架**

* beego (国内、初学者)

* gin (轻量级)

* echo (轻量级)

* lris (重量级)

2. **微服务框架**

* go kit --- 更加灵活，定制化

* Istio --- 繁琐的大型微服务项目

3. **Tcp长链接框架**

4. **容器编排**

* Kubernetes(k8s) --- 推荐

* swarm

5. **服务发现**

* consul

6. **存储引擎**

* k/v存储 --- etcd

* 分布式存储 --- tidb

7. **静态建站**

* hugo

8. **中间件**

很多...

* 消息队列 nsq

* Tcp长链接框架(轻量级服务器) zinx

* 游戏服务器 Leaf

* RPC框架 gRPC

* redis集群 codis  

9.  **爬虫框架**

* go query

