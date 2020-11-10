package main

import "fmt"

func main() {
	grid := uint32(63)
	sum := getNumberOfWheat(grid)
	fmt.Printf("前 %d 个棋格的麦粒总数为: %d \n", grid, sum)
}

// @title getNumberOfWheat
// @description 计算麦粒总数
// @param grid uint32 "格子数"
// @return sum uint64 "麦粒总数"
func getNumberOfWheat(grid uint32) (sum uint64) {
	sum = 0
	var numberOfWheatInGrid uint64 = 0

	//第一个格子的麦粒数
	numberOfWheatInGrid = 1

	sum = sum + numberOfWheatInGrid
	for i := uint32(2); i <= grid; i = i + 1 {
		numberOfWheatInGrid = numberOfWheatInGrid * 2
		sum = sum + numberOfWheatInGrid
	}

	return sum
}
