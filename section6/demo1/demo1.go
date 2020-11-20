package main

import "fmt"

func main() {
	var numbers = []int{10, 9, 20, 1, 3, 13, 19, 4}
	fmt.Println("归并排序，数字原始序列:")
	fmt.Print(numbers)
	numbers = mergeSort(numbers)
	fmt.Println("\n使用归并排序之后")
	fmt.Print(numbers)
}

func mergeSort(numbers []int) []int {
	if numbers == nil || len(numbers) == 0 {
		return make([]int, 0)
	}
	if len(numbers) == 1 {
		return numbers
	}
	mid := len(numbers) / 2
	left := numbers[0:mid]
	right := numbers[mid:len(numbers)]
	left = mergeSort(left)
	right = mergeSort(right)
	mergeResult := merge(left, right)
	return mergeResult
}

func merge(left []int, right []int) []int {
	var mergeResult []int = make([]int, len(left)+len(right))
	k := 0
	var i, j int
	for i, j = 0, 0; i < len(left) && j < len(right); {

		if left[i] <= right[j] {
			mergeResult[k] = left[i]
			i++
		} else {
			mergeResult[k] = right[j]
			j++
		}
		k++
	}

	if i == len(left) {
		for _, value := range right[j:] {
			mergeResult[k] = value
			k++
		}

	} else {
		for _, value := range left[i:] {
			mergeResult[k] = value
			k++
		}
	}

	return mergeResult
}
