package main

import (
	"fmt"
)

func main() {
	//常量定义之后，在整个声明周期中不能修改
	//const定以后可以不使用
	// 常量只支持:bool string 数字(整形，浮点数等等的)
	// 常量是编译过程中计算好的值
	//常量是在运行过程中进行配置得值(可用于配置文件)
	const pi float64 = 3.141592653
	fmt.Println("pi=", pi)
	// pi = 3.14

}
