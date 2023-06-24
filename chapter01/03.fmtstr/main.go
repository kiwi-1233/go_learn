package main

import (
	"fmt"
	"math"
)

func main() {
	//doc.golang.ltd //中文文档
	fmt.Printf("我的名字是%s\n", "小明")
	fmt.Printf("我的年龄是%q\n", "小强")
	fmt.Printf("我的年龄是%x\n", "小强")
	fmt.Printf("我的年龄是%X\n", "小强")

	const pi float64 = math.Pi
	var cir1R float64 = 2.34
	var cir2R float64 = 3.56
	var cir1 = pi * (cir1R * cir2R)
	fmt.Printf("园1的面积为:%f\n", cir1)
	var cir2 = pi * (cir2R * cir2R)
	fmt.Printf("园2的面积为:%f\n", cir2)
	fmt.Printf("两个圆的面积和为:%.3f\n", cir1+cir2)
}
