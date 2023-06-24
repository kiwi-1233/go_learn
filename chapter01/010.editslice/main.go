package main

import "fmt"

func main() {

	//string和byte可以互转，这样可以修改字符串（不是任意类型都可以转换）
	//byte类型对应的是ascii码表，所以string类型可以和byte互通而其他类型不可以
	//go 默认编码为utf8
	a := "你好"
	bytesSlice := []rune(a) //处理中文时要使用rune类型，因为byte类型支持的是ascii编码，无法处理中文，rune使用的是utf8
	fmt.Println(bytesSlice)
	fmt.Println("修改切片内容")
	bytesSlice[0] = 'H'
	a = string(bytesSlice)
	fmt.Println(a)

}
