package main

import (
	"fmt"
)

func main() {
	var number int = 8
	fmt.Printf("\n整数可以被分解为多个整数的乖积的组合")
	fmt.Printf("整数 %d 可以被分解为:\n", number)
	getResult(number)

	number = 0
	fmt.Printf("整数 %d 可以被分解为:\n", number)
	getResult(number)

	number = 4
	fmt.Printf("整数 %d 可以被分解为:\n", number)
	getResult(number)
}

// @title getResult
// @description 获取一个整数可以被分解为多个整数的乖积的组合
// para number int "输入的整数"
// returns none
func getResult(number int) {
	if number == 0 {
		fmt.Print("正数 0 不能分解乘积\n")
	}
	var divdeds []int = make([]int, 0)
	var result = make([]int, 0)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			divdeds = append(divdeds, i)
			//divdeds = append(divdeds, -i)
		}
	}
	split(divdeds, number, result)
}

// @title getResult
// @description 获取一个整数可以被分解为多个整数的乖积的组合
// para divdeds []int "可以被number整除的数字",number int "输入的整数",result []int "用于保存中间结果"
// returns none
func split(divdeds []int, number int, result []int) {

	//var initValue = 1
	if number == 1 || number == -1 {
		fmt.Print(result)
		fmt.Print("\n")
	} else if number == 0 {
		return
	}

	for i := 0; i < len(divdeds); i++ {
		containsInitValueOne := false
		containsInitValueMinusOne := false
		for k := 0; k < len(result); k++ {
			if 1 == result[k] {
				containsInitValueOne = true
			}
			if -1 == result[k] {
				containsInitValueMinusOne = true
			}
		}
		if containsInitValueOne && 1 == divdeds[i] {
			continue
		}
		if containsInitValueMinusOne && -1 == divdeds[i] {
			continue
		}

		newResult := make([]int, len(result))
		copy(newResult, result)
		newResult = append(newResult, divdeds[i])
		split(divdeds, number/divdeds[i], newResult)
	}
}
