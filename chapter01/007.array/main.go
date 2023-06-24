package main

import "fmt"

func main() {
	var ages [5]int = [5]int{1, 2, 3, 4, 5}
	var ages2 = [5]int{1, 2, 3, 4, 5}
	ages3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(ages)
	fmt.Println(ages2)
	fmt.Println(ages3)
	e := [...]int{} //不定长数组
	fmt.Println(e)
	//数组属于变量范畴，所以声明方式与变量类似
	//数组长度固定
	//数组类型固定，数组内的元素类型也固定
	//动态赋值==运行过程中数组元素的值可以改变
	array := [3]int{}
	array[0] = 1
	array[1] = 2
	array[2] = 33
	lenArray := len(array)
	fmt.Println(lenArray)

	for i, val := range array {
		fmt.Println(i, val)
	}

}
