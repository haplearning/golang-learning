package main

import "fmt"
import "math"

//变量

// var name = "asdf"  //函数外声明不能用短变量声明
//func main() {
//	var name string  //变量声明, 声明类型，默认初始化值为""
//	var age = 12  //变量声明，指定初始化值
//	var email, pwd = "test@test.com", 123
//	var isOk bool
//	var (  //批量声明
//		a = 1
//		b = 'b'
//		c = true
//		d float32
//	)
//
//	//m := 200  //短变量声明 :=
//	n := 20
//
//	fmt.Println(name, age, isOk)
//	fmt.Println(email, pwd)
//	fmt.Println(a, b, c, d)
//	fmt.Println(m, n)
//}

//匿名变量 _
//func foo() (int,string) {
//	return 10,"Q1mi"
//}
//
//func main(){
//	a, _ := foo()
//	_, b := foo()
//	fmt.Print(a,b)
//}

var name = "qimi"
var age int

// 常量
const (
	pi = 3.14
	//e = 2.7182
)
// 同时声明多个常量时，如果升略了值，则表示和上一行的值相同
//const (
//	n1 = 100
//	n2
//	n3
//)

//常量记数器 iota
//const (
//	n1 = iota  //初始为0
//	n2  //1
//	_  // 跳过某些值
//	n4  //3
//	n5 = 100 //100  插队
//	n6 = iota  //5
//)
//const n7 = iota
//
//const (
//	a, b = iota+1, iota+2  //1,2
//	c, d  //2,3
//	e, f  //3,4
//	g, h = iota, iota
//)

const (
	a1 = iota
	a2 = 100
	a3 = iota+1
	a4, a5, a6= iota+1, iota+1, iota+2 //44
	a7, a8, a9 //5 5 6
	//a3 = 3
	//a4 = iota+1

)
const  a10 = iota

const (
	_ = iota
	kb = 1<<(10*iota)
	mb = 1<<(10*iota)
	gb = 1<<(10*iota)
	tb = 1<<(10*iota)
	pb = 1<<(10*iota)
)


func main() {
	//fmt.Println(pi, e, name, age)
	//fmt.Println(n1, n2, n4, n5, n6, n7)
	//fmt.Println(a, b, c, d, e, f, g, h)
	//fmt.Println(a1, a2, a3, a4, a5, a6, a7,a8, a9,a10)
	//fmt.Println(kb, mb, gb, tb, pb)
	//fmt.Println(1<<10, 1<<20, 1<<30,1<<40,1<<50)
	//数字字面量语法
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a)  //10
	fmt.Printf("%b \n", a)  // 1010 占位符%b表示二进制

	var b int = 077
	fmt.Printf("%o \n", b)

	var c int = 0xff
	fmt.Printf("%x \n", c)
	fmt.Printf("%X \n", c)

	var d float64 = 3.1415926
	fmt.Println(math.MaxFloat64)
	fmt.Printf("%f \n", d)
	fmt.Printf("%.2f \n", d)

	var f complex64
	f = 1 + 2i
	fmt.Println(f)
}



