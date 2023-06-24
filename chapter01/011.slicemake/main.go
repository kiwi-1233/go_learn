package main

import "fmt"

func main() {
	// make 的长度参数，会初始化元素，为默认值0
	a := make([]int, 0)
	fmt.Println(len(a), cap(a))
	a = append(a, 1)
	fmt.Println(len(a), a)

	b := make([]int, 1)
	fmt.Println(len(b), cap(b))

	b = append(b, 1)
	fmt.Println(len(b), b)
	//初始化长度0 容量1的切片，go编译器在初始化时并不会准备好容量，
	//在容量扩展时底层是在不断地切换空间的，所以推荐在初始化时设定好容量
	c := make([]int, 0, 1)
	fmt.Println(len(c), cap(c))

	c = append(c, 1)
	fmt.Println(len(c), c)

}
