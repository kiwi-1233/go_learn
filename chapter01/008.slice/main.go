package main

import "fmt"

func main() {

	//切片的定义,切片是一个长度不固定存储特定元素的数组
	var a []int //空切片的声明 当中括号里有长度声明就是数据
	//var b [0]int //这是空数组

	// 切片的操作基本与数组相同
	//切片的增删改查
	//增
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	a = append(a, 3)
	a = append(a, 3)
	a = append(a, 3)
	a = append(a, 3)
	fmt.Println(a)
	//删，删除本质上就是把一个数组切开，拿走一部分再拼回去
	a = append(a[:1], a[2:5]...)
	fmt.Println(a)
	//增
	//backup:=append(a[1:]) //这样写，内存地址是共用的，所以当a发生变化时，backup也会发生改变
	backup := append([]int{}, a[1:]...) //所以需要重新创建一个切片
	a = append(a[:1], 999)
	a = append(a, backup...)
	fmt.Println(a)
	//查
	//通过下标
	b := a[3]
	fmt.Println("b:", b)
	//for循环 遍历
	//for i := 0; i < len(a); i++ {
	//	fmt.Println("slice index:", i, "slice val:", a[i])
	//}
	//for i, v := range a {
	//	fmt.Println("slice index:", i, "slice val:", v)
	//}

	//a2 := a
	a2 := make([]int, 0, len(a))
	fmt.Println("a2原始", a2, "a2长度：", len(a2), "a2容量：", cap(a2))
	a2 = append(a2, a...)
	a2[len(a2)-1] = 100
	fmt.Println("a:", a)
	fmt.Println("a2:", a2)

}
