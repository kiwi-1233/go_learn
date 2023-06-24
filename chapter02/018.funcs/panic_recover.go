package main

import "fmt"

func panicAndRecover() {

	defer coverPanicUpgraded() //可以

	//defer func() {  // 不可以
	//	coverPanicUpgraded()
	//}()

	//defer coverPanic() //不可以

	var nameScore map[string]int = nil
	nameScore["key"] = 100

}

func coverPanic() {
	func() {
		if r := recover(); r != nil {
			fmt.Println("系统出现严重故障：", r)
		}
	}()
}

func coverPanicUpgraded() {
	if r := recover(); r != nil {
		fmt.Println("系统出现严重故障：", r)
	}
}
