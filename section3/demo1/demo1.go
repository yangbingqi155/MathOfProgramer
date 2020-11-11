package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	grid := uint32(63)

	start1 := time.Now()
	sum := getNumberOfWheat(grid)
	end1 := time.Now()
	delta1 := end1.Sub(start1)

	fmt.Printf("前 %d 个棋格的麦粒总数为: %d,耗时 %d ns \n", grid, sum, delta1.Nanoseconds())

	start2 := time.Now()
	sum = uint64(math.Pow(2, float64(grid))) - 1
	end2 := time.Now()
	delta2 := end2.Sub(start2)
	fmt.Printf("前 %d 个棋格的麦粒总数为: %d,耗时 %d ns \n", grid, sum, delta2.Nanoseconds())

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
