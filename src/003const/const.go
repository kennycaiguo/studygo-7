package main

import "fmt"

// 常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。
const pi = 3.14159
const e = 2.7182

const (
	day = "1998-5-1"
	c   = 299792458
)

// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	n1 = 100
	n2
	n3
	n4
)

// iota是go语言的常量计数器，只能在常量的表达式中使用。
// iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
const (
	m1 = iota //0
	m2        //1
	m3        //2
	m4        //3
)

//使用_跳过某些值
const (
	l1 = iota //0
	l2        //1
	_         //2
	l3        //3
)

//iota声明中间插队
const (
	i1 = iota //0
	i2 = 100  //1
	i3 = iota //2
	i4        //3
)
const i5 = iota //重置为0

// 定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

// 多个iota定义在一行
const (
	a, b = iota + 1, iota + 2 //0
	x, d                      //1
	y, f                      //2
)

func main() {
	fmt.Println(pi, e, day, c)
	fmt.Println(n1, n2, n3, n4)
	fmt.Println(m1, m2, m3, m4)
	fmt.Println(l1, l2, l3)
	fmt.Println(i1, i2, i3, i4, i5)
	fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println(a, b, x, d, y, f)

}
