package main

import "fmt"

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

//func main()  {
//	s1:="aasdfasdfasdf"
//	s2:="123"
//	s3 := "asd"+"123"  //字符串拼接
//	s4 := fmt.Sprintf(s1, s2)  //字符串拼接
//	fmt.Println(s1, s2, s3, s4)
//	fmt.Println(len(s1))  //字符串长度,utf-8 是变长字符集，英文标点占用1个字节，中文占用3个字节。
//	fmt.Println(s1[0])  //字符串索引,返回的是ascii字节码
//	fmt.Printf("%c", s1[0])  //字符串索引字符
//	fmt.Println("abc">"abcd")  //字符串比较 ==,>,<，比较机制是字符的对称比较
//	fmt.Println(strings.Compare(s1,s2))  //字符串比较，比较机制是字符的对称比较，返回值0：a==b,1:a>b,-1:a<b
//
//	fmt.Println(strings.Contains(s1, "as"))  //检测substr是否在s中
//	fmt.Println(strings.ContainsAny(s1, "ah"))  //检测chars中的任意字符是否出现在s中
//	fmt.Println(strings.ContainsAny(s1, "hj"))
//	fmt.Println(strings.ContainsRune(s1, 'a'))  //检测 rune 字符是否出现在s中
//	fmt.Println(strings.ContainsRune(s1, 97))  // 检测 rune 字符是否出现在s中，可用ascii码代替
//
//	fmt.Println(strings.Count(s1, "asd"))  //统计s中非重叠substr的数量，若substr为“”，则返回len(s)+1
//	fmt.Println(strings.Count(s1, ""))
//
//	fmt.Println(strings.EqualFold("asd","asD"))  //检测s和t在忽略大小写的情况下是否相等
//
//	fmt.Println(strings.Split(s1, "a"))  //使用sep切割s,返回字符串切片
//	fmt.Println(strings.SplitN("go-Js-JavaScript","-", 2))  //在sep分割s, 使用n限制分割的元素数量，n=0返回nil,n<0表示不限制
//	fmt.Println(strings.SplitN("go-Js-JavaScript","-", 0))
//	fmt.Println(strings.SplitN("go-Js-JavaScript","-", -1))
//	fmt.Println(strings.SplitAfter("go-Js-JavaScript","-"))  //在sep后分割字符串s,返回字符串切片
//	fmt.Println(strings.SplitAfterN("go-Js-JavaScript","-", 2))  //在sep后分割s,使用n限制分割的元素个数，n=0返回nil,n<0表示不限制
//	fmt.Println(strings.SplitAfterN("go-Js-JavaScript","-", 0))
//	fmt.Println(strings.SplitAfterN("go-Js-JavaScript","-", -1))
//
//
//
//	fmt.Println(strings.Fields("a s d f"))  //返回空格分割的字符串s, 等价于strings.Split("a s d f"," ")
//	ff := func(c rune) bool {
//		return strings.ContainsRune(",|/", c)
//	}
//	fmt.Println(strings.FieldsFunc("go,python,c++/c,Js|JavaScript", ff))  //使用函数确定分隔符，来分割字符串
//
//	fmt.Println(strings.HasPrefix("asdfasdf","as")) //检测s是否已字符串prefix作为前缀
//	fmt.Println(strings.HasSuffix("asdfasdf","df")) //检测s是否已字符串suffix作为后缀
//
//	fmt.Println(strings.Index("asdfasdf","d"))  //返回substr在s中第一次出现的索引位置，若没有则返回-1
//	fmt.Println(strings.Index("asdfasdf","g"))
//	fmt.Println(strings.IndexAny("asdfasdf","lmnd")) //返回chars中任意字符在s中第一次出现的索引位置，若没有出现则返回-1
//	fmt.Println(strings.IndexByte("asdfasdf", 'h'))  //返回byte字符c在s中第一次出现的位置，若没有则返回-1
//	fi := func(c rune) bool {
//		return strings.ContainsRune(",|/",c)
//	}
//	fmt.Println(strings.IndexFunc("go,python,c++/c,Js|JavaScript", fi))  //返回s中第一次满足函数f的rune字符的索引位置
//	fmt.Println(strings.IndexRune("asdfasdf", 's'))  //返回rune字符r在s中第一次出现的索引位置，若没有则返回-1
//	fmt.Println(strings.LastIndex("asdfasdf","sd"))  //返回substr在s中最后一次出现的索引位置，若没有则返回-1
//	fmt.Println(strings.LastIndexAny("asdfasdf","agh"))  //返回chars任意字符在s中最后一次出现的索引位置，若没有则返回-1
//	fmt.Println(strings.LastIndexByte("asdfasdf",'d'))  //返回byte字符c在s中最后一次出现的索引位置，若没有则返回-1
//	fl := func(c rune) bool{
//		return strings.ContainsRune("./|", c)
//	}
//	fmt.Println(strings.LastIndexFunc("go,python,c++/c,Js|JavaScript",fl))  //返回s中字后一次满足函数fl的rune字符的索引位置，若没有则返回-1。
//	fm := func(c rune) rune{
//		if strings.ContainsRune(",|/", c){
//			return '-'
//		} else {
//			return c
//		}
//	}
//	fmt.Println(strings.Map(fm,"go,Js|JavaScript"))  //返回s中每个字符经过映射函数mapping处理之后的字符
//
//	fmt.Println(strings.Repeat("ps-", 3))  //返回s重复count次的字符串
//	fmt.Println(strings.Replace("asdfasdfasdf", "s","S",2)) //在s中使用new替换old,n限定替换次数，n为负数表示没有限制
//	fmt.Println(strings.Replace("asdfasdfasdf", "s","S",-1))
//
//
//	ss := []string{"go","hank","python","php"}
//	fmt.Println(strings.Join(ss, "-"))  //使用分隔符sep连接字符串切片a
//
//	fmt.Println(strings.Title("hello go !"))  //返回 Title 化的字符串(首字符大写)
//	fmt.Println(strings.ToTitle("hello go !"))  //返回所有字符 Title 化的字符串
//	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))  //使用特定的规则将全部字符s都Title化
//
//
//	fmt.Println(strings.ToLower("Hello Go !"))  //s小写转换
//	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))  //使用特定的规则转换s到小写
//	fmt.Println(strings.ToUpper("hello go !"))  //s大写转换
//	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))  //使用特定的规则转换s到大写
//
//	fmt.Println(strings.Trim(" user name "," "))  //去除s两端的特定字符cutset
//	ft:= func(c rune) bool{
//		return strings.ContainsRune(",/|",c)
//	}
//	fmt.Println(strings.TrimFunc("|/user name,/", ft))  //去除s两端满足函数ft的字符
//	fmt.Println(strings.TrimLeft(" user name ", " "))  //去除s左端的特定字符cutset
//	ftl:= func(c rune) bool{
//		return strings.ContainsRune(",/|",c)
//	}
//	fmt.Println(strings.TrimLeftFunc("|/user name,/", ftl))  //去除s左端的满足ftl的字符
//
//	fmt.Println(strings.TrimRight(" user name ", " "))  //去除s右端的特定字符cutset
//	ftr:= func(c rune) bool{
//		return strings.ContainsRune(",/|",c)
//	}
//	fmt.Println(strings.TrimRightFunc("|/user name,/", ftr))  //去除s右端的满足ftr的字符
//	fmt.Println(strings.TrimPrefix("asdf_asdf","asdf_"))  //去除s的前缀prefix
//	fmt.Println(strings.TrimSuffix("asdf_asdf","_asdf"))  //去除s的前缀suffix
//	fmt.Println(strings.TrimSpace(" \t\nhello go\n\t\r\n"))  //去除s两端的空白字符
//}

//遍历字符串
//func main(){
//	s := "hello沙河"
//	fmt.Println(len(s))
//	for i:=0; i<len(s);i++{
//		fmt.Println(i)
//		fmt.Printf("%v(%c)", s[i],s[i])
//	}
//	fmt.Println()
//	for i, r:= range s{
//		fmt.Println(i)
//		fmt.Printf("%v(%c) ", r, r)
//	}
//	fmt.Println()
//}

func main() {
	s1 := "big"

	byteS1:=[]byte(s1)
	byteS1[0]='p'
	fmt.Println(string(byteS1))

	s2:="白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}