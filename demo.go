package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
	"math/rand"
	"time"
	"sort"
)

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

//const (
//	a1 = iota
//	a2 = 100
//	a3 = iota+1
//	a4, a5, a6= iota+1, iota+1, iota+2 //44
//	a7, a8, a9 //5 5 6
//	//a3 = 3
//	//a4 = iota+1
//
//)
//const  a10 = iota
//
//const (
//	_ = iota
//	kb = 1<<(10*iota)
//	mb = 1<<(10*iota)
//	gb = 1<<(10*iota)
//	tb = 1<<(10*iota)
//	pb = 1<<(10*iota)
//)


//func main() {
//	//fmt.Println(pi, e, name, age)
//	//fmt.Println(n1, n2, n4, n5, n6, n7)
//	//fmt.Println(a, b, c, d, e, f, g, h)
//	//fmt.Println(a1, a2, a3, a4, a5, a6, a7,a8, a9,a10)
//	//fmt.Println(kb, mb, gb, tb, pb)
//	//fmt.Println(1<<10, 1<<20, 1<<30,1<<40,1<<50)
//	//数字字面量语法
//	// 十进制
//	var a int = 10
//	fmt.Printf("%d \n", a)  //10
//	fmt.Printf("%b \n", a)  // 1010 占位符%b表示二进制
//
//	var b int = 077
//	fmt.Printf("%o \n", b)
//
//	var c int = 0xff
//	fmt.Printf("%x \n", c)
//	fmt.Printf("%X \n", c)
//
//	var d float64 = 3.1415926
//	fmt.Println(math.MaxFloat64)
//	fmt.Printf("%f \n", d)
//	fmt.Printf("%.2f \n", d)
//
//	var f complex64
//	f = 1 + 2i
//	fmt.Println(f)
//
//	var g bool =true
//	fmt.Println(g)
//
//	var h = "asdf"  //只能双引号
//	fmt.Println(h)
//
//	fmt.Println("str := \"c\\code\\lesson1\\go.exe\"")  //’\‘转义字符串
//
//	s1 := `diyihang
//dierhang\\
//disanhang
//`
//	fmt.Println(s1)
//}

func str_funcs()  {
	s1:="aasdfasdfasdf"
	s2:="123"
	s3 := "asd"+"123"  //字符串拼接
	s4 := fmt.Sprintf(s1, s2)  //字符串拼接
	fmt.Println(s1, s2, s3, s4)
	fmt.Println(len(s1))  //字符串长度,utf-8 是变长字符集，英文标点占用1个字节，中文占用3个字节。
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
}

//遍历字符串
func traversal_str(){
	s := "hello沙河"
	fmt.Println(len(s))
	for i:=0; i<len(s);i++{
		fmt.Println(i)
		fmt.Printf("%v(%c)", s[i],s[i])
	}
	fmt.Println()
	for i, r:= range s{
		fmt.Println(i)
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func type_trans() {
	s1 := "big"

	byteS1:=[]byte(s1)
	byteS1[0]='p'
	fmt.Println(s1, string(byteS1))

	s2:="白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(s2, string(runeS2))

	var a,b = 3,4
	var c int
	c = int(math.Sqrt(float64(a*a +b*b)))
	fmt.Print(c)

	d:=1
	e:=1.2
	f:=true
	g:="aaa"

	fmt.Printf("%d(%t)\n",d, d)
	fmt.Printf("%d(%t)\n",e, e)
	fmt.Printf("%d(%t)\n",f, f)
	fmt.Printf("%d(%t)\n",g, g)

	s := "hello沙河小王子"
	count:=0
	runeS := []rune(s)
	for i:=0;i<len(runeS);i++{
		if len(string(runeS[i]))==3 {
			count += 1
		}
	}
	fmt.Printf("汉字数量：%d", count)
}

//运算符
func oprater_demo() {
	//算术运算符：+,-,*,/,=,%,
	//关系运算符：>,>=,<,<=,==,!=
	//逻辑运算符：&&(and),||(or),!(非)
	//位运算符：&(位与)，|(位或)，^(异或)，>>(右移n位，*2**n)，<<(左移n位，/2**n)
	//赋值运算符：=，+=，-=，*=，/=，%=，<<=，>>=，&=，|=，^=

	s := [5]int{1, 2, 3, 1, 2}
	fmt.Println(s[0]^s[1]^s[2]^s[3]^s[4])
}

//流程控制
//if else 分支结构
func ifdemo1 () {

	score:=65
	if score >= 90 {
		fmt.Println("A")
	}else if score>=80 {
		fmt.Println("B")
	}else{
		fmt.Println("C")
	}
	fmt.Println(score)
}

func ifdemo2 () {
	//if else 分支结构

	if score:=65; score >= 90 { //此时socre为局部变量
		fmt.Println("A")
	}else if score>=80 {
		fmt.Println("B")
	}else{
		fmt.Println("C")
	}
	//fmt.Println(score)
}

func fordemo()  {
	//语法：
	//for 初始语句;条件表达式;结束语句{
	//}
	for i:=0;i<10;i++{
		fmt.Println(i)
	}

	i:=0
	for ;i<10;i++{ //省略初始语句（需提前定义）
		fmt.Println(i)
	}

	i=0
	for i<10 {  //等价于while,省略初始语句、结束语句（需提前定义初始语句，在结构体中定义结束语句）
		fmt.Println(i)
		i++
	}
}

func forrangedemo(){
	s:=[]string{"a","b","c","d"}
	for i,v:=range s{
		fmt.Println(i,v)
	}
}

func switchdemo(){
	finger:=3
	switch finger {
	case 1:
		fmt.Println("a")
	case 2,3,4: //case分支可以有多个值
		fmt.Println("b")
	default:  //只能有一个default
		fmt.Println("无效的输入")
	}
	
	age :=30
	switch  {  //分支用表达式时，switch后无需再跟变量
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
		  //fallthrough语法可以执行满足条件case的下一个case
	case age > 60:
		fmt.Println("好好享受吧")
		fallthrough
	default:
		fmt.Println("活着真好")
	}
}

func gotodemo1(){
	var breakflag bool
	for i:=0;i<10;i++{
		for j:=0;j<10;j++{
			if j==2{
				breakflag=true
				break
			}
			fmt.Printf("%v-%v\n",i,j)
		}
		if breakflag{
			break
		}
	}
}

func gotodemo2()  {
	for i:=0;i<10;i++{
		for j:=0;j<10;j++{
			if j==2{
				goto breakflag
			}
			fmt.Printf("%v-%v\n",i,j)
		}
	}
	breakflag:
}

func breakdemo()  {
	breakdemo:
	for i:=0;i<10;i++{
		for j:=0;j<10;j++{
			if i==2&&j==2{
				break breakdemo
			}
			fmt.Println(i,j)
		}
	}
}

func continuedemo(){
	forloop:
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			if i==2&&j==2{
				continue forloop
			}
			fmt.Println(i,j)
		}
	}
}

func multtable(){
	for i:=1;i<10;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%v*%v=%v ",j,i,i*j)
		}
		fmt.Println()
	}
}

func revemulttale(){
	for i:=9;i>0;i--{
		for j:=1;j<=i;j++{
			fmt.Printf("%v*%v=%v ",j,i,i*j)
		}
		fmt.Println()
	}
}

//数组
func arrdemo(){
	//数组定义
	var a [3]int  //var 变量名 [元素数量]T
	var b [5]int  //长度固定，[3]int 和[5]int 是不同的类型,默认为初始值
	fmt.Println(a,b)
	//初始化
	var c [3] int  //默认初始值为零值
	var d = [3]int {1, 2}   //使用指定的值初始化
	var e = [3]string{"a","b","c"}
	var f = [3]string{"a","b"}
	fmt.Println(c,d,e,f[2])
	//自动确定数组长度
	var g = [...]int{1,2}
	fmt.Printf("%T",g)
	//指定索引值
	var h = [...]int{0:3,2:2}
	fmt.Println(h, len(h))
	fmt.Printf("%T",h)

	//数组遍历
	for i:=0;i<len(e);i++{
		fmt.Println(e[i])
	}
	for _,v := range d{
		fmt.Println(v)
	}
	//多维数组
	var i =[3][2]int{
		{1,2},{3,4},{5,6},
	}
	fmt.Println(i)
	fmt.Println(i[2][1])  //支持索引取值
	var j = [...][2]int{  //多维数组只能在外层让编译器推导长度
		{7,8},{9,10},
	}
	fmt.Printf("%T--%v",j,j)
	fmt.Println(j[0][1])


	//二维数组的遍历
	for _,v1 :=range i{
		for _,v2 :=range v1{
			fmt.Println(v2)
		}
	}

	k:=[]int{1,2,3}
	fmt.Println(len(k))
	k=append(k, 4)
	fmt.Println(k)


	fmt.Printf("%T",0)
	m:=[]int{}
	fmt.Println(m)

	//数组是值类型，赋值和传参会复制整个数组，改变副本的值，原数据不会改变

}

func modifyArr1(x [3]int)  {
	x[0]=100
}
func modifyArr2(x [][2]int){
	x[2][0]=100
}
func modifydemo(){
	a:=[...]int{10,20,30}
	modifyArr1(a)
	b:=a
	b[1]=200
	fmt.Println(a)

	c:=[][2]int{
		{1,2},{3,4},{5,6},
	}
	modifyArr2(c)
	fmt.Println(c)
}

//切片
func slice_demo(){
	//切片声明：var name []T
	var a []string  //声明切片
	var b = []int{}   //声明切片并初始化
	var c =[]bool{}  //声明切片并初始化
	//var d =[]bool{false,true}  //声明切片并初始化
	fmt.Println(a, b, c) //[] [] [flase,true]
	fmt.Println(a==nil)  //true
	fmt.Println(b==nil)  //false
	fmt.Println(c==nil)  //false
	//fmt.Println(c==d)  //切片是引用类型，不支持直接比较，只能和nil比较

	e:=[5]int{1,2,3,4,5}
	s:=e[1:3]  //s:=a[low:high] 简单表达式
	//e[:3]  等价于 e[0:3]
	//e[1:]  等价于 e[1:len(e)]
	//e[:]  等价于 e[0:len(e)]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n",s,len(s),cap(s))  //s:[2 3] len(s):2 cap(s):4
	s1 :=s[3:4]  //切片再切片时，索引上限时cap(s) 而不是 len(s)
	fmt.Printf("s1:%v len(s1):%v cap(s):%v\n", s1, len(s1), cap(s1))  //s1:[5] len(s1):1 cap(s):1

	t:=e[1:3:5]  //t:=e[low:high:max]  容量为max-low
	t1:=e[3:5]  //只有low 能省略
	fmt.Printf("t:%v len(t):%v cap(t):%v\n",t,len(t),cap(t))  //t:[2 3] len(t):2 cap(t):4
	fmt.Printf("t1:%v len(t1):%v cap(t1):%v\n",t1,len(t1),cap(t1))  //t1:[4 5] len(t1):2 cap(t1):2

	//使用make()构造切片: make([]T,size,cap)
	f:=make([]int, 2, 10)
	fmt.Printf("f:%v len(f):%v cap(f):%v\n",f,len(f),cap(f)) //f:[0 0] len(f):2 cap(f):10
	//切片本质：对底层数组的封装，包含以下三个信息：1)底层数组的指针 2)切片长度len 3)切片的容量cap
	var g =[]int{}
	fmt.Printf("type:%T g:%v len(g):%v cap(g):%v isnil:%v\n",g, g,len(g),cap(g), g==nil)

	//判断切片是否为空需用 len(s)==0 而不是 s==nil
	fmt.Printf("a:%v len(a):%v a==nil?%v\n", a,len(a),a==nil)
	fmt.Printf("a:%v len(a):%v a==nil?%v\n", b,len(b),b==nil)  //初始化虽len(s)=0 但 !=nil
	fmt.Printf("a:%v len(a):%v a==nil?%v\n", c,len(c),c==nil)  //初始化虽len(s)=0 但 !=nil

	//切片为引用类型，拷贝后变量共享底层数组，一个变化会影响另一个切片
	s2 := make([]int, 3)
	s3 :=s2
	s3[1]=100
	fmt.Println(s2)  //[0 100 0]
	fmt.Println(s3)  //[0 100 0]

	//切片遍历
	s4:=[]int{1,2,3,4}
	for i:=0;i<len(s4);i++{
		fmt.Println(i, s4[i])
	}

	for idx,v :=range s4{
		fmt.Println(idx, v)
	}

	//添加方法 append()
	var s5 []int   //var 声明的零值切片可直接再append()使用，无需初始化
	s5 = append(s5,1) //[1]
	s5 = append(s5, 2,3,4)  //[1 2 3 4]
	s6 :=[]int{5,6,7}  //[1 2 3 4 5 6 7]
	s5 = append(s5, s6...)
	fmt.Println(s5)

	//append()添加元素和切片扩容，每次扩容后都是扩容前的2倍
	var numslice []int
	for i:=0;i<10;i++{
		numslice = append(numslice,i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", numslice, len(numslice), cap(numslice),numslice)
	}

	//使用copy(destSlice, srcSlice []T) 将src切片元素复制到dest中
	h:=[]int{1,2,3,4,5}
	j:=h
	fmt.Println(h,j)  //[1 2 3 4 5] [1 2 3 4 5]
	j[0]=1000  //切片是引用类型，h、j都指向同一块内存空间
	fmt.Println(h, j)  //[1000 2 3 4 5] [1000 2 3 4 5]
	k:=[]int{1,2,3}
	copy(k, h)  //
	k[1]=2000
	fmt.Println(h, k)

	//删除切片中的元素没有专门的方法，可以通过切片特性删除
	l:=[]int{30,31,32,33,34}
	//删除索引为2的元素
	l=append(l[:2], l[3:]...)
	fmt.Println(l)
}

func homework(){
	var a=make([]string,5,10)
	for i:=0;i<10;i++{
		a=append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}

func map_demo(){
	//定义map  map[KeyType]ValueType
	//使用make 创建map：make(map[KeyType]ValueType, [cap])
	scoreMap:=make(map[string]int, 8)

	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	//声明时填充元素
	userinfo:=map[string]string{
		"username":"沙河小王子",
		"password":"12345",
	}
	fmt.Println(userinfo)

	//查看键是否存在 value,ok:=map[key]
	v, ok := scoreMap["小明"]  //key存在, ok为True, v为对应值; 否则ok为false, v值类型零值
	if ok{
		fmt.Println(v)
	}else{
		fmt.Println("查无此人")
	}
	
	//map遍历 for...range
	for k,v:= range scoreMap{
		fmt.Println(k,v)
	}
	//只遍历key
	for k:=range scoreMap{
		fmt.Println(k)
	}

	//delete 删除键值对: delete(map, key)
	delete(scoreMap,"小明")
	fmt.Println(scoreMap)

	//按指定顺序遍历map
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoremap = make(map[string]int, 200)
	for i:=0;i<100;i++ {
		key:=fmt.Sprintf("stu%02d", i)  //生stu开头的字符串
		value:=rand.Intn(100)  //生成0-99的随机整数
		scoremap[key]=value
	}
	//取出map中的所有key存入切片keys
	var keys=make([]string,0,20)
	for key:= range scoremap{
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key 遍历map
	// for _,key:=range keys{
	// 	fmt.Println(key, scoremap[key])
	// }

	// 元素类型为map的切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice{
		fmt.Sprintf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	//对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"]="小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value:=range mapSlice{
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	//值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key:="中国"
	value, ok:=sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

func main()  {
	//ifdemo1()
	//ifdemo2()
	//fordemo()
	//forrangedemo()
	//switchdemo()
	//gotodemo2()
	//continuedemo()
	//multtable()
	//revemulttale()
	//arrdemo()
	//modifydemo()
	// slice_demo()
	// homework()
	map_demo()

}
