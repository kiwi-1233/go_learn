package main

import "fmt"

var counter int

func calcSum(nums ...int) (sum int) {
	for _, item := range nums {
		sum += item
	}
	counter++
	return
}

func showUsedTimes() {
	fmt.Println("used:", counter)
}

func genImprovementFunc() func(percentage float64) {
	base := 1000.0
	return func(percentage float64) {
		base = base * (1 + percentage)
		fmt.Println(base)
	}
}

func closerMain() {
	imp := genImprovementFunc()
	imp(0.1)
	imp(0.1)
	imp(0.1)
	imp(0.1)

}

func count() func() {
	i := 0
	return func() {
		i++
		fmt.Println(i)
	}
}
