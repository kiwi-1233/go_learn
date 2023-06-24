package main

import (
	"fmt"
)

func main() {
	var buy_apple_num int = 6
	var seeMallan bool = false
	if seeMallan {
		buy_apple_num = 1

	} else {
		buy_apple_num = 7
	}
	fmt.Printf("买了%d个苹果\n", buy_apple_num)

}
