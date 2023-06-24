package main

import "fmt"

func main() {
	//翻转1维数组
	array1 := [...]int{1, 3, 5, 7, 8, 9, 0}

	var endIndex = len(array1) - 1
	for i := 0; i < len(array1)/2; i++ {
		startValue := array1[i]
		array1[i] = array1[endIndex]
		array1[endIndex] = startValue
		endIndex--

	}
	fmt.Println(array1)

}
