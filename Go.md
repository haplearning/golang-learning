# 一、Go 语言之变量与常量

### 1、标识符与关键字

#### 1）标识符

在编程语言中，标识符就是程序员定义的具有特殊意义的词，比如变量名、常量名、函数名等等。Go 语言中标识符由字母、数字和`_` （下划线）组成，并且只能以字母和`_` 开头。例如：`abc`，`_`，`_123`，`a124`。

#### 2）关键字

关键字是指编程语言中预先定义好的具有特殊含义的标识符。 关键字和保留字都不建议用作变量名

Go 语言中有 25 个关键字：

```go
    break        default      func         interface    select
    case         defer        go           map          struct
    chan         else         goto         package      switch
    const        fallthrough  if           range        type
    continue     for          import       return       var
```

此外，Go 语言中还有 37 个保留字

```go
    Constants:    true  false  iota  nil

        Types:    int  int8  int16  int32  int64  
                  uint  uint8  uint16  uint32  uint64  uintptr
                  float32  float64  complex128  complex64
                  bool  byte  rune  string  error

    Functions:   make  len  cap  new  append  copy  close  delete
                 complex  real  imag
                 panic  recover
```



### 2、变量

#### 1）变量来历

程序运行过程中的数据都是保存在内存中的，我们想要在代码中操作某个数据就需要去内存中找到这个变量，但如果我们直接在代码中通过内存地址去操作变量的话，代码的可读性会变得非常差而且还容易出错，所以我们就利用变量将这个数据的内存地址保存起来，以后直接通过这个变量就能找到内存上对应的数据了。

#### 2）变量类型

变量（Variable）的功能是存储数据。不同的变量保存的数据类型可能会不一样。经过半个多世纪的发展，编程语言已经基本形成了一套固定的类型，常见变量的数据类型有：整型、浮点型、布尔型等。

#### 3）变量声明

Go语言中的每一个变量都有自己的类型，并且变量必须经过声明才能开始使用

##### 标准声明

Go 语言的变量声明格式为：

```go
var 变量名 变量值
```

变量声明以关键字 `var` 开头，变量类型放再变量的后面，行尾无需分号。例如：

```go
var name string
var age int
var isOk bool
```

##### 批量声明

每声明一个变量就需要写 `var` 关键字会比较繁琐，go语言中还支持批量变量声明：

```go
var (
    a string
    b int
    c bool
    d float32
)
```

##### 变量的初始化

Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的默认值，例如： 

- 整型和浮点型变量的默认值为`0`。 
- 字符串变量的默认值为`空字符串`。 
- 布尔型变量默认为`false`。 
- 切片、函数、指针变量的默认为`nil`。

当然我们也可在声明变量的时候为其指定初始值。变量初始化的标准格式如下：

```go
var 变量名 类型 = 表达式
```

例如：

```go
var name string = 'Q1mi'
var age int = 18
```

或者一次初始化多个变量

```go
var name, age = 'Q1mi', 19 
```

##### 短变量声明

在函数内部，可以使用更简略的 `:=` 方式声明并初始化变量。

```go
package main

import "fmt"

// 全局变量 m
var m = 100

func main {
    n := 20
    m := 200 //此处声明局部变量
    fmt.Println(m, n)
}
```

##### 匿名变量

在使用多重赋值时，如果想要忽略某个值，可以使用`匿名变量（anonymous variable）`。 匿名变量用一个下划线`_`表示，例如

```go
package main

func foo() (int, string){
    return 10, "qimi"
}

func main(){
    x, _ := foo()
    _, y := foo()
    fmt.Println("x=", x)
    fmt.Println("y=", y)
}
```

匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。 (在`Lua`等编程语言里，匿名变量也被叫做哑元变量。)

#### 注意事项：

1. 函数外的每个语句都必须以关键字开始（var、const、func等）
2. `:=`不能使用在函数外。
3. `_`多用于占位，表示忽略值。

### 3、常量

#### 1）常量声明

相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。 常量的声明和变量声明非常类似，只是把`var`换成了`const`，常量在定义的时候必须赋值

```go
const pi = 3.14
const e = 2.7182
```

声明了`pi`和`e`这两个常量之后，在整个程序运行期间它们的值都不能再发生变化了。

多个常量也可以一起声明：

```go
const (
	pi = 3.14
    e = 2.7182
)
```

`const` 同时声明多个常量时，如果省略了值则表示和上面一行的值相同。 例如：

```go
const (
	n1 = 100
    n2
    n3
)
```

上面示例中，常量`n1`、`n2`、`n3`的值都是100。

#### 2）`iota` 常量计数器

`iota`是 go 语言的常量计数器，只能在常量的表达式中使用。

`iota` 在 const 关键字出现时将被重置为 0。const 中每新增一行常量声明将使`iota`计数一次( iota 可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。

举个例子：

```go
const (
		n1 = iota //0
		n2        //1
		n3        //2
		n4        //3
	)
```

##### 几种常见的 iota 示例：

使用 `_` 跳过某些值

```go
const (
		n1 = iota //0
		n2        //1
		_
		n4        //3
	)

```

`iota`声明中间插队

```go
const (
	n1 = iota  //0
    n2 = 100  //100
    n3 = iota  //2
    n4  //3
)
const n5 = 0
```

定义数量级 （这里的`<<`表示左移操作，`1<<10`表示将1的二进制表示向左移10位，也就是由`1`变成了`10000000000`，也就是十进制的1024。同理`2<<2`表示将2的二进制表示向左移2位，也就是由`10`变成了`1000`，也就是十进制的8。）

```go
const(
	_ = iota
    KB = 1 << (10 * iota)
    MB = 1 << (10 * iota)
    GB = 1 << (10 * iota)
    TB = 1 << (10 * iota)
    PB = 1 << (10 * iota)
)
```

多个 `iota` 定义一行

```go
const(
	a, b = iota + 1, iota + 2 //1, 2
    c, d  //2, 3
    e, f  //3, 4
)
```

# 二、Go 语言之基本数据类型

Go语言中有丰富的数据类型，除了基本的整型、浮点型、布尔型、字符串外，还有数组、切片、结构体、函数、map、通道（channel）等。Go 语言的基本类型和其他语言大同小异。

### 1、基本数据类型

#### 1）整型

整型分为以下两个大类： 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64

其中，`uint8` 就是我们熟知的`byte` 型，`int16`对应C语言中的`short`型，`int64`对应C语言中的`long`型。

|  类型  | 描述                                                         |
| :----: | :----------------------------------------------------------- |
| uint8  | 无符号 8位整型 (0 到 255)                                    |
| uint16 | 无符号 16位整型 (0 到 65535)                                 |
| uint32 | 无符号 32位整型 (0 到 4294967295)                            |
| uint64 | 无符号 64位整型 (0 到 18446744073709551615)                  |
|  int8  | 有符号 8位整型 (-128 到 127)                                 |
| int16  | 有符号 16位整型 (-32768 到 32767)                            |
| int32  | 有符号 32位整型 (-2147483648 到 2147483647)                  |
| int64  | 有符号 64位整型 (-9223372036854775808 到 9223372036854775807) |

##### 特殊整型

|  类型   |                          描述                          |
| :-----: | :----------------------------------------------------: |
|  uint   | 32位操作系统上就是`uint32`，64位操作系统上就是`uint64` |
|   int   |  32位操作系统上就是`int32`，64位操作系统上就是`int64`  |
| uintptr |              无符号整型，用于存放一个指针              |

**注意：** 在使用`int`和 `uint`类型时，不能假定它是32位或64位的整型，而是考虑`int` 和`uint` 可能在不同平台上的差异。

**注意事项**：获取对象的长度的内建`len()`函数返回的长度可以根据不同平台的字节长度进行变化。实际使用中，切片或 map 的元素数量等都可以用`int`来表示。在涉及到二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用`int`和 `uint`。

##### 数字字面量语法（Number literals syntax）

Go1.13版本之后引入了数字字面量语法，这样便于开发者以二进制、八进制或十六进制浮点数的格式定义数字，例如：

`v := 0b00101101`， 代表二进制的 101101，相当于十进制的 45。 `v := 0o377`，代表八进制的 377，相当于十进制的 255。 `v := 0x1p-2`，代表十六进制的 1 除以 2²，也就是 0.25。

而且还允许我们用 `_` 来分隔数字，比如说： `v := 123_456` 表示 v 的值等于 123456。

我们可以借助fmt函数来将一个整数以不同进制形式展示。

```go
package main
 
import "fmt"
 
func main(){
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a)  // 10
	fmt.Printf("%b \n", a)  // 1010  占位符%b表示二进制
 
	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b)  // 77
 
	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c)  // ff
	fmt.Printf("%X \n", c)  // FF
}
```

#### 2）浮点数

Go语言支持两种浮点型数：`float32`和`float64`。这两种浮点型数据格式遵循`IEEE 754`标准： `float32` 的浮点数的最大范围约为 `3.4e+38`，可以使用常量定义：`math.MaxFloat32`。 `float64` 的浮点数的最大范围约为 `1.8e+308`，可以使用一个常量定义：`math.MaxFloat64`。

打印浮点数时，可以使用`fmt`包配合动词`%f`，代码如下：

```go
package main
import (
        "fmt"
        "math"
)
func main() {
        fmt.Printf("%f\n", math.Pi)
        fmt.Printf("%.2f\n", math.Pi)
}
```

#### 3）复数

`complex64` 和`complex128`

```go
var c1 complex64
c1 = 1 + 2i
var c2 complex128
c2 = 2 + 3i
fmt.Println(c1)
fmt.Println(c2)
```

复数有实部和虚部，`complex64` 的实部和虚部为32位，`complex128` 的实部和虚部为64位。

#### 4）布尔值

Go语言中以`bool`类型进行声明布尔型数据，布尔型数据只有`true（真）`和`false（假）`两个值。

**注意：**

1. 布尔类型变量的默认值为`false`。
2. Go 语言中不允许将整型强制转换为布尔型.
3. 布尔型无法参与数值运算，也无法与其他类型进行转换。

#### 5）字符串

Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。 Go 语言里的字符串的内部实现使用`UTF-8`编码。 字符串的值为`双引号(")`中的内容，可以在Go语言的源码中直接添加非ASCII码字符，例如：

```go
s1 := "hello"
s2 := "你好"
```

##### 字符串转义符

Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表所示。

| 转义符 |                含义                |
| :----: | :--------------------------------: |
|  `\r`  |         回车符（返回行首）         |
|  `\n`  | 换行符（直接跳到下一行的同列位置） |
|  `\t`  |               制表符               |
|  `\'`  |               单引号               |
|  `\"`  |               双引号               |
|  `\\`  |               反斜杠               |

举个例子，我们要打印一个Windows平台下的一个文件路径：

```go
package main
import (
    "fmt"
)
func main() {
    fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
}
```

##### 多行字符串

Go语言中要定义一个多行字符串时，就必须使用`反引号`字符：

```go
s1 := `第一行
第二行
第三行
`
fmt.Println(s1)
```

反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。

##### 字符串的常用操作

|                方法                 |      介绍      |
| :---------------------------------: | :------------: |
|              len(str)               |     求长度     |
|           +或fmt.Sprintf            |   拼接字符串   |
|            strings.Split            |      分割      |
|          strings.contains           |  判断是否包含  |
| strings.HasPrefix,strings.HasSuffix | 前缀/后缀判断  |
| strings.Index(),strings.LastIndex() | 子串出现的位置 |
| strings.Join(a[]string, sep string) |    join操作    |

```go
s1:="aasdfasdfasdf"
s2:="123"
s3 := "asd"+"123"  //字符串拼接
s4 := fmt.Sprintf(s1, s2)  //字符串拼接
fmt.Println(s1, s2, s3, s4)
fmt.Println(len(s1))  //字符串长度,utf-8是变长字符集，英文标点占用1个字节，中文占3个字节
fmt.Println(s1[0])  //字符串索引,返回的是ascii字节码
fmt.Printf("%c", s1[0])  //字符串索引字符
fmt.Println("abc">"abcd")  //字符串比较 ==,>,<，比较机制是字符的对称比较
fmt.Println(strings.Compare(s1,s2))  //字符串比较，比较机制是字符的对称比较，返回值0：a==b,1:a>b,-1:a<b

fmt.Println(strings.Contains(s1, "as"))  //检测substr是否在s中
fmt.Println(strings.ContainsAny(s1, "ah"))  //检测chars中的任意字符是否出现在s中
fmt.Println(strings.ContainsAny(s1, "hj"))
fmt.Println(strings.ContainsRune(s1, 'a'))  //检测 rune 字符是否出现在s中
fmt.Println(strings.ContainsRune(s1, 97))  // 检测 rune 字符是否出现在s中，可用ascii码代替

fmt.Println(strings.Count(s1, "asd"))  //统计s中非重叠substr的数量，若substr为“”，则返回len(s)+1
fmt.Println(strings.Count(s1, ""))

fmt.Println(strings.EqualFold("asd","asD"))  //检测s和t在忽略大小写的情况下是否相等

fmt.Println(strings.Split(s1, "a"))  //使用sep切割s,返回字符串切片
fmt.Println(strings.SplitN("go-Js-JavaScript","-", 2))  //在sep分割s, 使用n限制分割的元素数量，n=0返回nil,n<0表示不限制
fmt.Println(strings.SplitN("go-Js-JavaScript","-", 0))
fmt.Println(strings.SplitN("go-Js-JavaScript","-", -1))
fmt.Println(strings.SplitAfter("go-Js-JavaScript","-"))  //在sep后分割字符串s,返回字符串切片
fmt.Println(strings.SplitAfterN("go-Js-JavaScript","-", 2))  //在sep后分割s,使用n限制分割的元素个数，n=0返回nil,n<0表示不限制
fmt.Println(strings.SplitAfterN("go-Js-JavaScript","-", 0))
fmt.Println(strings.SplitAfterN("go-Js-JavaScript","-", -1))



fmt.Println(strings.Fields("a s d f"))  //返回空格分割的字符串s, 等价于strings.Split("a s d f"," ")
ff := func(c rune) bool {
    return strings.ContainsRune(",|/", c)
}
fmt.Println(strings.FieldsFunc("go,python,c++/c,Js|JavaScript", ff))  //使用函数确定分隔符，来分割字符串

fmt.Println(strings.HasPrefix("asdfasdf","as")) //检测s是否已字符串prefix作为前缀
fmt.Println(strings.HasSuffix("asdfasdf","df")) //检测s是否已字符串suffix作为后缀

fmt.Println(strings.Index("asdfasdf","d"))  //返回substr在s中第一次出现的索引位置，若没有则返回-1
fmt.Println(strings.Index("asdfasdf","g"))
fmt.Println(strings.IndexAny("asdfasdf","lmnd")) //返回chars中任意字符在s中第一次出现的索引位置，若没有出现则返回-1
fmt.Println(strings.IndexByte("asdfasdf", 'h'))  //返回byte字符c在s中第一次出现的位置，若没有则返回-1
fi := func(c rune) bool {
    return strings.ContainsRune(",|/",c)
}
fmt.Println(strings.IndexFunc("go,python,c++/c,Js|JavaScript", fi))  //返回s中第一次满足函数f的rune字符的索引位置
fmt.Println(strings.IndexRune("asdfasdf", 's'))  //返回rune字符r在s中第一次出现的索引位置，若没有则返回-1
fmt.Println(strings.LastIndex("asdfasdf","sd"))  //返回substr在s中最后一次出现的索引位置，若没有则返回-1
fmt.Println(strings.LastIndexAny("asdfasdf","agh"))  //返回chars任意字符在s中最后一次出现的索引位置，若没有则返回-1
fmt.Println(strings.LastIndexByte("asdfasdf",'d'))  //返回byte字符c在s中最后一次出现的索引位置，若没有则返回-1
fl := func(c rune) bool{
    return strings.ContainsRune("./|", c)
}
fmt.Println(strings.LastIndexFunc("go,python,c++/c,Js|JavaScript",fl))  //返回s中字后一次满足函数fl的rune字符的索引位置，若没有则返回-1。
fm := func(c rune) rune{
    if strings.ContainsRune(",|/", c){
        return '-'
    } else {
        return c
    }
}
fmt.Println(strings.Map(fm,"go,Js|JavaScript"))  //返回s中每个字符经过映射函数mapping处理之后的字符

fmt.Println(strings.Repeat("ps-", 3))  //返回s重复count次的字符串
fmt.Println(strings.Replace("asdfasdfasdf", "s","S",2)) //在s中使用new替换old,n限定替换次数，n为负数表示没有限制
fmt.Println(strings.Replace("asdfasdfasdf", "s","S",-1))


ss := []string{"go","hank","python","php"}
fmt.Println(strings.Join(ss, "-"))  //使用分隔符sep连接字符串切片a

fmt.Println(strings.Title("hello go !"))  //返回 Title 化的字符串(首字符大写)
fmt.Println(strings.ToTitle("hello go !"))  //返回所有字符 Title 化的字符串
fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))  //使用特定的规则将全部字符s都Title化


fmt.Println(strings.ToLower("Hello Go !"))  //s小写转换
fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))  //使用特定的规则转换s到小写
fmt.Println(strings.ToUpper("hello go !"))  //s大写转换
fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))  //使用特定的规则转换s到大写

fmt.Println(strings.Trim(" user name "," "))  //去除s两端的特定字符cutset
ft:= func(c rune) bool{
    return strings.ContainsRune(",/|",c)
}
fmt.Println(strings.TrimFunc("|/user name,/", ft))  //去除s两端满足函数ft的字符
fmt.Println(strings.TrimLeft(" user name ", " "))  //去除s左端的特定字符cutset
ftl:= func(c rune) bool{
    return strings.ContainsRune(",/|",c)
}
fmt.Println(strings.TrimLeftFunc("|/user name,/", ftl))  //去除s左端的满足ftl的字符

fmt.Println(strings.TrimRight(" user name ", " "))  //去除s右端的特定字符cutset
ftr:= func(c rune) bool{
    return strings.ContainsRune(",/|",c)
}
fmt.Println(strings.TrimRightFunc("|/user name,/", ftr))  //去除s右端的满足ftr的字符
fmt.Println(strings.TrimPrefix("asdf_asdf","asdf_"))  //去除s的前缀prefix
fmt.Println(strings.TrimSuffix("asdf_asdf","_asdf"))  //去除s的前缀suffix
fmt.Println(strings.TrimSpace(" \t\nhello go\n\t\r\n"))  //去除s两端的空白字符
```

#### 6）byte 和 rune 类型

组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。

字符用单引号（‘）包裹起来，如：

```go
var a := '中'
var b := 'x'
```

Go 语言的字符有以下两种：

- unit8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符
- rune 类型，代表一个 UTF-8 字符

当需要处理中文、日文或者其他复合字符时，则需要用到`rune`类型。`rune`类型实际是一个`int32`。

Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。

```go
func main(){
	s := "hello沙河"
	for i:=0; i<len(s);i++{
		fmt.Printf("%v(%c)", s[i],s[i])
	}
	fmt.Println()
	for _, r:= range s{
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
```

输出：

```go
104(h)101(e)108(l)108(l)111(o)230(æ)178(²)153()230(æ)178(²)179(³)
104(h) 101(e) 108(l) 108(l) 111(o) 27801(沙) 27827(河) 
```

因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。

字符串底层是一个byte数组，所以可以和`[]byte`类型相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。

##### 修改字符串

要修改字符串，需要先将其转换成`[]rune`或`[]byte`，完成后再转换为`string`。无论哪种转换，都会重新分配内存，并复制字节数组。