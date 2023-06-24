package main

import "fmt"

func main() {
	//多维数组
	infos := [3][3]string{ //长度确定难以维护
		[3]string{"1", "2", "3"},
		[3]string{"2", "2", "3"},
		[3]string{"2", "2", "3"},
	}
	for _, val := range infos {
		fmt.Println(val)
	}

	infos2 := [...][3]string{ //支持动态添加
		[3]string{"1", "2", "3"},
		[3]string{"2", "2", "3"},
		[3]string{"2", "2", "3"},
		[3]string{"2", "2", "3"},
		[3]string{"2", "2", "3"},
		[3]string{"2", "2", "3"},
		[3]string{"2", "2", "3"},
		[3]string{"2", "2", "3"},
	}
	for _, val := range infos2 {
		fmt.Println(val)
	}
	for _, val1 := range infos2 {
		fmt.Println(val1)
		for _, val2 := range val1 {
			fmt.Println(val2)
		}
	}

}
