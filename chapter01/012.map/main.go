package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Student struct {
	name    string
	math    int
	english int
	chinese int
}

func average(s Student) float64 {
	return float64(s.math+s.english+s.chinese) / 3
}

func main() {
	//map的声明
	//var map1 map[string]int = nil
	//map2 := map[string]int{}
	//map3 := map[string]int{"小强": 1, "小李": 2, "小张": 3}
	//
	////map的基本操作
	////增
	//map2["哈哈"] = 1
	//map2["嘿嘿"] = 2
	//fmt.Println("map1:", map2)
	////删
	//delete(map2, "哈哈")
	////map在删除不存在的key时不会报错
	//delete(map2, "呵呵")
	//fmt.Println("map1:", map2)
	////改
	//map2["嘿嘿"] = 3
	//fmt.Println("map1:", map2)
	////查
	//val := map2["嘿嘿"]
	//fmt.Println("val:", val)
	////map在查询不存在的key的时候不会报错,
	//val2 := map2["哈哈"]
	//fmt.Println("val2:", val2)
	//
	////遍历
	//for k, v := range map3 {
	//	fmt.Printf("k:%s=v:%d\n", k, v)
	//}
	//
	////map使用的注意事项
	////map不用实例化就可以读取，删除
	//fmt.Println("未实例化时读取map")
	//val3 := map1["小红"]
	//fmt.Println("val3=", val3)
	//fmt.Println("未实例化时删除map")
	//delete(map1, "小红")
	////map没有实例化的时候不可写
	////map1["小红"] = 333  //报错：panic: assignment to entry in nil map
	////重复删除map不会引发异常
	////如何确定是否在map中取到了值
	////查询map的某一个值时，会返回两个值，查询到的value和是否为真实的值，可以避免一些错误
	//fakeVal, no := map3["小明"]
	//fmt.Println(fakeVal, ">>>>>>>>", no)
	//map3["小明"] = 88
	//realVal, ok := map3["小明"]
	//fmt.Println(realVal, ">>>>>>>>", ok)

	//map练习

	//用map管理20个人的分数，算出所有人的平均分，高分在前
	//1. 创建二十个人的分数map
	//scoreMap := map[string][]int{}
	//avgMap := map[string]int{}
	//var sortSlice []int
	//
	s1 := rand.NewSource(time.Now().Unix()) //用指定值创建一个随机数种子
	r1 := rand.New(s1)
	//
	////fmt.Println(scoreMap)
	////计算平均分
	//for k, v := range scoreMap {
	//	avgScore := (v[0] + v[1] + v[2]) / 3
	//	avgMap[k] = avgScore
	//	sortSlice = append(sortSlice, avgScore)
	//
	//}
	//fmt.Println(avgMap)

	//fmt.Println(avgMap)
	//TODO:排序并输出

	students := make(map[string]Student)
	//students["张三"] = Student{"张三", 90, 80, 85}
	//students["李四"] = Student{"李四", 95, 75, 90}
	for i := 1; i <= 20; i++ {
		name := fmt.Sprintf("姓名%d", i)
		score1 := r1.Intn(100) + 1
		score2 := r1.Intn(100) + 1
		score3 := r1.Intn(100) + 1
		students[name] = Student{name, score1, score2, score3}

	}

	//func average(s Student) float64 {
	//	return float64(s.math + s.english + s.chinese) / 3
	//}

	scores := make([]struct {
		name  string
		score float64
	}, len(students))

	i := 0
	for k, v := range students {
		scores[i].name = k
		scores[i].score = average(v)
		i++
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].score > scores[j].score // 按照降序排列
	})

	prevScore := -1.0 // 初始值设为-1.0表示没有上一次输出
	for _, s := range scores {
		if s.score == prevScore { // 如果与上一次输出相同，则在同一行输出，并用逗号隔开
			fmt.Printf(", %s %.2f", s.name, s.score)
		} else { // 如果与上一次输出不同，则换行输出，并更新prevScore为当前元素的平均分
			fmt.Printf("\n%s %.2f", s.name, s.score)
			prevScore = s.score
		}
	}

}
