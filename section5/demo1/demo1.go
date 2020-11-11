package main

import "fmt"

func main() {
	fmt.Println("\n求解给定的一个总金额，有1，2，5，10 这4种面额的纸币，总共有多少种奖赏方法(排列组合) \n")
	var money = []int{1, 2, 5, 10}
	var amount = 10
	result := make([]int, 0)
	fmt.Printf("输入总金额: %d\n", amount)
	get(money, amount, result)
}

// @title get
// @description 获取一个给定的金额有多少种奖赏方式，使用不同面额的纸币
// @para money []int "不同面额的纸币" , amount int "给定金额" ,result []int用于存放 每个组合纸币的面额
// @return none
func get(money []int, amount int, result []int) {

	if amount == 0 {

		fmt.Print(result)
		fmt.Print("\n")
		return
	} else if amount < 0 {
		return
	} else {
		for i := 0; i < len(money); i = i + 1 {
			newResult := make([]int, len(result))
			copy(newResult, result)
			newResult = append(newResult, money[i])
			get(money, amount-money[i], newResult)
		}
	}

}
