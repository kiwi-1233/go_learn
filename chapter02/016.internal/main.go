package main

import (
	"fmt"
)

func main() {
	testPanic()

}

func testPanic() {
	arr2 := make([]int, 3, 4)

	//复制arr2 覆盖 arr3
	fmt.Println(arr2[2])
	fmt.Println(constructHellos("1", "2"))
	getScoresOfStudent("1")
}

func constructHello(name string) string {
	return fmt.Sprintf("hello %s", name)
}

func constructHellos(name ...string) (ret string) {
	ret = fmt.Sprintf("hello %s", name)
	return
}
func getScoresOfStudent(name string) (int, int, int, int) {
	chinese, english, math, nature := 90, 82, 77, 81
	return chinese, english, math, nature
}
