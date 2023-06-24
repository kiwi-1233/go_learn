package main

import "fmt"

func main() {

	var money int
	if 0 < money && money <= 20 {
		fmt.Println("点外卖")
	} else if 20 < money && money <= 200 {
		fmt.Println("下馆子")
	} else if 200 < money && money <= 2000 {
		fmt.Println("去米其林")
	} else if 2000 < money && money <= 20000 {
		fmt.Println("出国玩儿一圈")
	} else {
		fmt.Println("梦里啥都有")
	}

}
