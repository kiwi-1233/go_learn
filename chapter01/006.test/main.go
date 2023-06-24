package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		height float32
		weight float32
		gender int
		age    int
	)
	for {

		fmt.Println("请输入您的年龄:")
		fmt.Scanln(&age)

		fmt.Println("请输入您的性别(女:0,男:1)")
		fmt.Scanln(&gender)
		if gender != 0 && gender != 1 {
			fmt.Println("输入的性别不合法")
			continue
		}
		fmt.Println("请输入您的身高(米):")
		fmt.Scanln(&height)
		fmt.Println("请输入您的体重(公斤)")
		fmt.Scanln(&weight)
		fmt.Println(height, weight, age, gender)
		if gender == 1 {
			fmt.Println("您的性别为:男性")
		} else if gender == 2 {
			fmt.Println("您的性别为:女性")
		} else {
			fmt.Println("性别未知")
		}
		fmt.Printf("您的年龄为%d岁\n", age)
		bmi := weight / (height * height)

		if 0 < bmi && bmi <= 18.5 {
			fmt.Printf("您的BMI指数为:%.2f 偏瘦\n", bmi)
		} else if 18.5 < bmi && bmi <= 24.0 {
			fmt.Printf("您的BMI指数为:%.2f 标准\n", bmi)
		} else if 24.0 < bmi && bmi <= 28 {
			fmt.Printf("您的BMI指数为:%.2f 肥胖\n", bmi)
		} else if bmi > 28.0 {
			fmt.Printf("您的BMI指数为:%.2f 严重肥胖\n", bmi)
		} else {
			fmt.Println("BMI值不合法")
		}

		bodyFat := 1.2*bmi + 0.23*float32(age) - 5.4 - 10.8*float32(gender)

		if 0 < bodyFat && bodyFat <= 11 {
			fmt.Printf("您的体脂率为:%.2f非常瘦\n", bodyFat)
		} else if 11 < bodyFat && bodyFat <= 17.0 {
			fmt.Printf("您的体脂率为:%.2f偏瘦\n", bodyFat)
		} else if 17.0 < bodyFat && bodyFat <= 22.0 {
			fmt.Printf("您的体脂率为:%.2f标准\n", bodyFat)
		} else if 22.0 < bodyFat && bodyFat <= 27.0 {
			fmt.Printf("您的体脂率为:%.2f肥胖\n", bodyFat)
		} else if bodyFat > 27 {
			fmt.Printf("您的体脂率为:%.2f严重肥胖\n", bodyFat)
		} else {
			fmt.Println("体脂率不合法")
		}
		var isContinue string
		fmt.Println("是否继续录入(Y/N):")
		fmt.Scanln(&isContinue)

		if strings.ToUpper(isContinue) == "N" {
			break
		}
	}

}
