package main

import "fmt"

func main() {
	//for循环结构 for  变量;循环执行条件;循环次数{代码块}
	for i := 0; i < 3; i++ {
		fmt.Println("你好 golang")
	}
	fmt.Println("=========================")
	//for循环的简写,for结构的三个部分都可以省略 只写退出条件时，可以不写分号 其他情况做省略都需要写分号
	j := 0
	for ; j < 5; j++ {
		fmt.Println("你好 golang")

	}
	fmt.Println("=========================")
	k := 0
	for k < 5 {
		fmt.Println("你好 golang")
		k++
	}
	fmt.Println("=========================")
	//三个条件都可以省略，但循环次数和退出条件都必须存在于循环体中
	l := 0
	for {
		fmt.Println("你好 golang")
		l++
		if l >= 5 {
			break
		}
	}
	fmt.Println("=========================")
	m := 0
	for {

		m++
		if m >= 50 {
			break
		}
		if m%2 == 0 {
			fmt.Printf("continue:%d\n", m)
			continue
		} else {
			fmt.Println(m)
		}

	}
}
