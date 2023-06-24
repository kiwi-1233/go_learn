package main

import "fmt"

func main() {
	var money float32 = 200
	var busy bool
	switch money {
	case 20:
		fmt.Println("20")
	case 200:
		fmt.Println("200")
		if busy {
			break
		} else {
			fmt.Println("吃零食")
		}
		fallthrough // 直接进入下一个分支，无需条件，例如状态机中会使用
	case 2000:
		fmt.Println("2000")
	case 20000:
		fmt.Println("20000")
		//default:
		//	fmt.Println(money)
	}
	// 分支表达式可以替代if表达式，并且有if表达式不具备的功能
	//1 可以随时中止
	//2. 并入其他分支继续执行
	//3. 原生支持类型判断
	// 4 . 判断类型的同时赋值新的变量，且类型自动匹配
	//5. case 可以是多条件的

}
