package main

import "fmt"

func main() {
	//位运算

	// & 位与:两个数同一位都是1 才为1  有一个为0 则为0

	// | 位或：两个数同一位，有一个是1 就是1

	// ^ 异或:两个数同一位，不一样的才是1
	a, b := 4, 3
	fmt.Println("a & b=", a&b)
	fmt.Println("b & a=", b&a)
	fmt.Println("a | b=", a|b)
	fmt.Println("b | a=", b|a)
	fmt.Println("a ^ b=", a^b)
	fmt.Println("b ^ a=", b^a)
	//异或交换律，两个数字进行异或，无论谁在前，结果都一样

	// <<  >> 位移运算
	arr := []int{4, 3, 4, 5, 6, 7, 3, 5, 6}
	result := -1
	for _, item := range arr {
		fmt.Println("item=", item)
		if result < 0 {
			result = item
			fmt.Println("result=", result)
		} else {
			result = result ^ item
			fmt.Println("result=", result)

		}

	}
	fmt.Println(result)
}
