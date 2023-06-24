package main

import "fmt"

func main() {
	// 计算过程中要注意数据类型长度，否则会溢出
	var a, b int8 = 30, 10

	fmt.Println("a + b=", a+b)
	fmt.Println("a - b=", a-b)
	fmt.Println("a * b=", a*b)
	fmt.Println("a / b=", a/b) // 除法结果会丢弃小数
	fmt.Println("a % b=", a%b) // 取余只能用证书计算

	//比较运算
	// && || !  == > < >= <= !=

}
