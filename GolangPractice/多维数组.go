package main

import "fmt"

func main() {
	animals := [][]string{}

	row1 := []string{"1", "2", "3"}
	row2 := []string{"3", "4"}
	row3 := []string{"5", "6", "7"}
	// 使用 append() 函数将一维数组添加到二维数组中
	animals = append(animals, row1)
	animals = append(animals, row2)
	animals = append(animals, row3)

	// 循环输出
	for i := range animals {
		fmt.Printf("Row: %v\n", i)
		fmt.Println(animals[i])
	}

}
