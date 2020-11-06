package main

import "fmt"

//Int32Max int32 max value
const Int32Max = int32(^uint32(0) >> 1)

//Int32Min int32 min value
const Int32Min = ^Int32Max

func main() {
	fmt.Println("自定义十进制转二进制(通过二进制的移位和逻辑与操作运算来实现)")
	s1 := IntToBinary(Int32Max)
	fmt.Printf("十进制数 %d 转成二进制后是: %s \n", Int32Max, s1)

	s2 := IntToBinary(Int32Min)
	fmt.Printf("十进制数 %d 转成二进制后是: %s \n", Int32Min, s2)
}

//IntToBinary int covert to binary
func IntToBinary(number int32) string {
	tempn := number
	if number < 0 {
		//求补码，因为负数是以补码的方式存储的
		number = ((-number) ^ Int32Max) + 1
	}
	s := ""
	var i int32
	i = 0
	for i = 31; i >= 0; i = i - 1 {
		//处理最高位
		if i == 31 {
			if tempn >= 0 {
				s = s + "0"
			} else {
				s = s + "1"
			}
		} else {
			//逻辑与位运算
			bit := int32(0)
			if (int32(1)<<i)&number > 0 {
				bit = int32(1)
			}
			s = s + fmt.Sprintf("%d", bit)
		}

	}
	return s
}

//Reverse reverse string
func Reverse(s string) string {
	temps := []rune(s)
	for i, j := 0, len(temps)-1; i < j; i, j = i+1, j-1 {
		temps[i], temps[j] = temps[j], temps[i]
	}
	return string(temps)
}
