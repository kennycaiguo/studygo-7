package main

import "fmt"

func main() {
	ifDemo1(22)
	ifDemo2()
	forDemo1()
	forDemo2()
	forDemo3()
	switchDemo1()
	switchDemo2()
	switchDemo3()
	switchDemo4()
	gotoDemo1()
	gotoDemo2()
	breakDemo1()
	continueDemo()
}

func ifDemo1(age int) {
	if age < 20 {
		fmt.Println("年轻真好")
	} else if age < 60 {
		fmt.Println("好好工作吧")
	} else {
		fmt.Println("好好享受晚年生活吧")
	}
}

//可以在if判断语句中直接声明变量，变量作用域只在if判断语句内
func ifDemo2() {
	var grade byte
	if score := 65; score > 90 {
		grade = 'A'
	} else if score > 75 {
		grade = 'B'
	} else if score > 60 {
		grade = 'C'
	} else {
		grade = 'D'
	}
	score := grade //这里再声明一个score也不会冲突
	fmt.Printf("%c\n", score)
}

// 条件表达式返回true时循环体不停地进行循环，直到条件表达式返回false时自动退出循环。
func forDemo1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// for循环的初始语句可以被忽略，但是初始语句后的分号必须要写
func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// for循环的初始语句和结束语句都可以省略
func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
}

// 这种写法类似于其他编程语言中的while，在while后添加一个条件表达式，满足条件表达式时持续循环，否则结束循环。

// 使用switch语句可方便地对大量的值进行条件判断。
// Go语言规定每个switch只能有一个default分支。
func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}

// 一个分支可以有多个值，多个case值中间使用英文逗号分隔。
func switchDemo2() {
	switch n := 7; n { //也可以在语句内声明变量
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

// 分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。
func switchDemo3() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}

// fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
func switchDemo4() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

// goto(跳转到指定标签)
// goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。
// Go语言中使用goto语句能简化一些代码的实现过程。 例如双层嵌套的for循环要退出时：
func gotoDemo1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}
}

// 使用goto语句能简化代码：
func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}

// break(跳出循环)
// break语句可以结束for、switch和select的代码块。
// break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的for、switch和 select的代码块上。
func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}

//continue(继续下次循环)
// continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用。
// 在 continue语句后添加标签时，表示开始标签对应的循环。例如：
func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
