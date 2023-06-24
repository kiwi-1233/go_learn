package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		name   [3]string
		height [3]float32
		weight [3]float32
		gender [3]int
		age    [3]int

		bmi1      [3]float32
		bodyFat1  [3]float32
		genderStr string

		bmiAll     float32
		bodyFatAll float32

		index int
	)
	for {
		fmt.Println("请输入您的姓名")
		fmt.Scanln(&name[index])

		fmt.Println("请输入您的年龄:")
		fmt.Scanln(&age[index])

		fmt.Println("请输入您的性别(男:1,女:0)")
		fmt.Scanln(&gender[index])
		switch gender[index] {
		case 0:
			genderStr = "女"
		case 1:
			genderStr = "男"
		default:
			fmt.Println("性别非法，请重新录入")
			continue
		}

		fmt.Println("请输入您的身高(米):")
		fmt.Scanln(&height[index])

		fmt.Println("请输入您的体重(公斤)")
		fmt.Scanln(&weight[index])

		fmt.Println(height[index], weight[index], age[index], gender[index])

		//fmt.Printf("姓名:%s\n", name[index])
		//fmt.Printf("性别:%s\n", genderStr)
		//fmt.Printf("年龄:%d岁\n", age[index])
		//fmt.Printf("身高:%.2f米\n", height[index])
		//fmt.Printf("体重:%.2f公斤\n", weight[index])

		bmi := weight[index] / (height[index] * height[index])
		if 0 < bmi && bmi <= 18.5 {
			fmt.Printf("您的BMI指数为:%.2f 偏瘦\n", bmi)
		} else if 18.5 < bmi && bmi <= 24.0 {
			fmt.Printf("您的BMI指数为:%.2f 标准\n", bmi)
		} else if 24.0 < bmi && bmi <= 28 {
			fmt.Printf("您的BMI指数为:%.2f 肥胖\n", bmi)
		} else if bmi > 28.0 {
			fmt.Printf("您的BMI指数为:%.2f 严重肥胖\n", bmi)
		} else {
			fmt.Println("BMI值不合法请重新录入")
			continue
		}

		bodyFat := 1.2*bmi + 0.23*float32(age[index]) - 5.4 - 10.8*float32(gender[index])
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
			fmt.Println("体脂率不合法请重新录入")
			continue
		}

		var isContinue string
		fmt.Println("是否继续录入(Y/N):")
		fmt.Scanln(&isContinue)
		bmi1[index] = bmi
		bodyFat1[index] = bodyFat
		index += 1

		if strings.ToUpper(isContinue) == "N" {
			break
		}

	}
	for i := 0; i <= index; i++ {
		fmt.Printf("姓名:%s,性别:%s,年龄:%d,身高:%.2f,体重:%.2f,bmi:%.2f,体脂率:%.2f\n", name[i], genderStr, age[i], height[i], weight[i], bmi1[i], bodyFat1[i])
		bmiAll += bmi1[i]
		bodyFatAll += bodyFat1[i]
	}

	var bmiAvg float32 = bmiAll / 3
	var bodyFatAvg float32 = bodyFatAll / 3
	fmt.Printf("当前录入人数:%d人\n", index+1)
	fmt.Printf("平均bmi为：%.2f\n", bmiAvg)
	fmt.Printf("平均体脂率为为：%.2f\n", bodyFatAvg)

	var isQuit string
	fmt.Println("按Q退出:")
	fmt.Scanln(&isQuit)

}
