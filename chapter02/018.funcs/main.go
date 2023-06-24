package main

import (
	"fmt"
	"time"
)

// 全局变量
var tall float64
var weight float64

func main() {
	panicAndRecover()

	fmt.Println("===========================")
	numCount := count()
	numCount()
	numCount()
	numCount()
	numCount()
	fmt.Println("===========================")
	openFile()
	fmt.Println("===========================")
	closerMain()
	fmt.Println("===========================")
	deferGuess()
	fmt.Println("===========================")
	time.Sleep(10 * time.Second)
	//closer call
	fmt.Println(calcSum(1, 2, 3, 4, 5))
	fmt.Println(calcSum(1, 2, 3, 4, 5))
	fmt.Println(calcSum(1, 2, 3, 4, 5))
	fmt.Println(calcSum(1, 2, 3, 4, 5))
	showUsedTimes()

	// 上级作用域的成员对下级作用域可见
	tall = 1.88
	fmt.Printf("全局的tall：%v\n", tall)
	weight = 100
	fmt.Printf("全局的weight：%v\n", weight)

	//下级作用域可定义与上级作用域同名的成员
	tall, weight := 1.70, 170.0
	fmt.Printf("局部的tall：%v\n", tall)
	fmt.Printf("局部的weight：%v\n", weight)

	//作用域可嵌套
	calculatorAdd := func(a, b float64) float64 {
		return a + b
	}
	result := calculatorAdd(1, 2)

	fmt.Println(result)
	//可以单独用花括号表示作用域
	{
		person1Tall := 1.81
		person1Weight := 100.0
		calculatorAdd(person1Tall, person1Weight)

	}
	{
		person1Tall := 1.81
		person1Weight := 100.0
		calculatorAdd(person1Tall, person1Weight)
	}
	//sampleSubdomain()
	//guess(1, 100)

}

func calcBMI() float64 {
	return 0
}

// 作用域操作时遵循就近原则
func sampleSubdomain() {
	name := "小明"
	fmt.Println("父作用域", name)
	{
		name := "小明"
		name = "小小明"
		fmt.Println("子作用域", name)
	}
	fmt.Println("父作用域", name)
}

func sampleSubdomainIf() {
	// v 在整个条件语句的代码块中都有效，在条件语句外无效,分支语句也一样
	if v := calcBMI(); v == 0 {
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}
}

func sampleSubdomainFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

//猜数字

func guess(left, right uint) {
	fmt.Println("left:", left, "right:", right)
	if left == right {
		fmt.Println("错误")
		return
	}
	guessed := (left + right) / 2
	fmt.Printf("是%v吗？\n", guessed)
	var getInput string
	fmt.Println("高了：1，低了：0，对了：9")
	fmt.Scanln(&getInput)
	switch getInput {
	case "1":
		guess(left, guessed-1)
	case "0":
		guess(guessed+1, right)
	case "9":
		fmt.Println("数字是:", guessed)

	}
}
