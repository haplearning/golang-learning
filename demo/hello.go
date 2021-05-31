package main

import (
	"flag"
	"fmt"
)

/*简单实现 hello world
func main()  {
	fmt.Print("hello world !")
}

//for 循环
func main()  {
	var s, sep string
	for i:=1; i<len(os.Args[1:]); i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Print(s)
}


//range 返回 idx,value
func main() {
	for _, arg := range os.Args{
		fmt.Print(arg)
	}
}


//while 循环
func main() {
	var i int
	for i<10{
		fmt.Print(i)
		i++
	}
}


//strings.Join() 字符串拼接
func main() {
	fmt.Print(strings.Join(os.Args," "))
}


func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Print(input.Text())
	print(input.Text())
}


func main() {
	counts := make(map[string]int)  //创建map键值对 string为键类型（通常），int为值类型
	input := bufio.NewScanner(os.Stdin)  //获取键盘输入
	for input.Scan(){  //读取下一行
		counts[input.Text()]++  //input.Text() 获取读到的内容
	}
	// 注意：忽略 input.Err() 中可能的错误
	for line, n := range counts{
		if n>1 {
			fmt.Print("%d\t%s\n", n, line)
		}
	}
}

// 声明：给实体命名，并设定其部分或全部属性
// 变量 var \常量 const \类型 type \函数 func
// 声明的通用形式：var name type = expression
// 短变量声明：name := expression
// 指针：x:=1 对应指针 p:=&x  通过指针访问内容 *p 通过指针修改内容 *p:=2
const boilingF = 212.0  //main 包级别的变量

func main() {
	var f = boilingF  //main函数的局部变量
	var c = (f-32)*5/9  //main函数的局部变量
	fmt.Printf("boiling point = %gF or %g\n", f, c)
}


func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g F = %g C\n", freezingF, fToc(freezingF))
	fmt.Printf("%g F = %g C\n", boilingF, fToc(boilingF))

}

func fToc(f float64) float64{
	return (f-32)*5/9
}
*/

func main() {
	//var s string
	//var i, j ,k int
	//var b, f, a = true, 2.3, "four"
	//fmt.Println(s, i, j, k, b, f, a)
	//x:=1  //定义变量
	//p:=&x  //指针
	//fmt.Println(x, p, *p)
	//fmt.Println(*p==x)
	//*p = 2
	//fmt.Println(x)
	//v :=1
	//incr(&v)
	//fmt.Println(incr(&v))
	//fmt.Println(v)
	//flag.Parse()
	//fmt.Print(strings.Join(flag.Args(), *sep))
	//fmt.Println(*n, sep)
	//if !*n{
	//	fmt.Println()
	//p := new(int)
	//fmt.Println(*p)
	//*p = 3
	//fmt.Println(*p)
	//p:=new(string)
	//q:=new(string)
	fmt.Println(delta(5,3))

}

func f() *int{
	v:=1
	return &v
}
func incr(p *int) int {
	*p++
	return *p
}

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

//func newInt() *int {
//	return new(int)
//}
//func newInt() *int {
//	var dummy int
//	return &dummy
//
//}
func delta(old, new int) int {
	return new - old
}

