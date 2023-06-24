package main

import "fmt"

func helloToTheOne(name string) string {
	return fmt.Sprintf("你好，%s", name)
}

func main() {
	retString := helloToTheOne("月之")
	fmt.Println(retString)
}
