package main

import (
	"fmt"
	"reflect"
)

func main() {
	var money interface{} = 10.32   //可以装任何类型
	switch newVar := money.(type) { //此时newVar是万能类型,在case之前类型不会确定，所以在case中计算时要注意
	case int:
		fmt.Println("int", newVar+2.0)
	case int64:
		fmt.Println("int64", newVar+22)

	case float32, float64:
		fmt.Println("float32", "newVar是小数")
	default:
		fmt.Println("unknow")
	}
	fmt.Println("结束", money, reflect.TypeOf(money))

}
