package main

import "fmt"

func main() {
	x1, y1 := getPointXYFromInput()
	x2, y2 := getPointXYFromInput()
	x3, y3 := getPointXYFromInput()
	x4, y4 := getPointXYFromInput()

	k1 := calculateK(y2, y1, x2, x1)
	k2 := calculateK(y4, y3, x4, x3)

	getResult(k1, k2)

}

func getResult(k1 float64, k2 float64) {
	fmt.Printf("直线1斜率为%.1f\n", k1)
	fmt.Printf("直线2斜率为%.1f\n", k2)
	if k1 == k2 {
		fmt.Println("平行")
	} else {
		fmt.Println("不平行")
	}
}

func getPointXYFromInput() (float64, float64) {
	var x, y float64
	fmt.Println("录入x")
	fmt.Scanln(&x)

	fmt.Println("录入y")
	fmt.Scanln(&y)
	return x, y

}

func calculateK(y2, y1, x2, x1 float64) float64 {
	k1 := (y2 - y1) / (x2 - x1)
	return k1
}
