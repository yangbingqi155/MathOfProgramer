package main

import (
	"fmt"
	"math"
)

func main() {
	var number uint64 = 4
	squreRoot := getSqureRoot(number, 0.00000000000001, 100000)
	if squreRoot == -1 {
		fmt.Printf("请输入大1的正整数\n")
	} else if squreRoot == -2 {
		fmt.Printf("未能找到解\n")
	} else {
		fmt.Printf("%d 的平方根是 %g ", number, squreRoot)
	}
}

// @title getSqureRoot
// @description 获取一个正整数的平方根
// @para n uint64 "给定的正整数" , deltaThreshold float64 "精度值" , maxTry uint32 "最大执行次数"
// @return squreRoot 返回正整数的平方根
func getSqureRoot(n uint64, deltaThreshold float64, maxTry uint32) (squreRoot float64) {
	if n < 1 {
		return -1
	}
	min := float64(1)
	max := float64(n)
	for i := uint32(0); i < maxTry; i = i + 1 {
		mid := min + (max-min)/float64(2)
		squre := mid * mid
		delta := math.Abs((squre / float64(n)) - 1)
		if delta <= deltaThreshold {
			return mid
		} else {
			if squre < float64(n) {
				min = mid
			} else if squre > float64(n) {
				max = mid
			}
		}
	}
	return -2

}
