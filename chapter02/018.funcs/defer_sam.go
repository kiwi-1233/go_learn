package main

import (
	"fmt"
	"os"
	"time"
)

func openFile() {
	fileName := "test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	defer file.Close()
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("close err", err)
			return
		}
	}()
	file = nil

}

func deferGuess() {
	startTime := time.Now()
	defer fmt.Println("duration：", time.Now().Sub(startTime))
	time.Sleep(5 * time.Second)
	fmt.Println("start time:", startTime)
}
