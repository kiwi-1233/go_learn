package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// go是强类型语言 声明和后续赋值，类型必须保持一致
	// go变量必须先定义后使用

	//变量定义的多样性
	var hello string = "hello,golang"
	fmt.Println("========================")
	//类型推导，go会猜测变量类型，
	var world = "world"
	fmt.Println(hello, world)
	var3 := 1.2333
	fmt.Println(var3)
	//类型推导的问题:猜测未必准确，所以一定要确认需要的类型
	a1 := "hello"
	fmt.Println(reflect.TypeOf(a1))
	a2 := 3
	fmt.Println(reflect.TypeOf(a2))
	a3 := 3.0
	fmt.Println(reflect.TypeOf(a3))
	a4 := &a3
	fmt.Println(reflect.TypeOf(a4))
	fmt.Println("========================")
	//一次定义多个变量
	var var4, var5 uint = 33, 44
	fmt.Println(var4, var5)
	//分行批量定义
	var (
		var6        = 7
		var7 int    = 8
		var8 string = "9"
	)
	fmt.Println(var6, var7, var8)

	//变量使用
	fmt.Println(var6 * var7)
	fmt.Println("========================")
	//进行算术运算时，变量类型必须一致
	var ho, ver float32 = 3, 4.22
	var sc = ho * ver
	fmt.Println(ho * ver)
	fmt.Println(sc)

	// 变量定义后必须使用

	//变量不初始化会被赋予默认值 int默认是0  string默认是空字符串
	fmt.Println("========================")
	var stringEmpty string
	var intZero int
	fmt.Println("string_empty:", stringEmpty, "int_zero:", intZero)

	//变量类型转换及代价，高类型转低类型会丢失精度
	fmt.Println("========================")
	var f1 float64 = 1.234
	fmt.Println(f1)
	var i1 int = int(f1)
	fmt.Println(i1)
	var i2 int = 23
	fmt.Println(i2)
	var f2 float64 = float64(i2)
	fmt.Println(f2)
	var a6 uint = math.MaxUint64
	var a7 int = int(a6)
	fmt.Println(a6, a7)

}
