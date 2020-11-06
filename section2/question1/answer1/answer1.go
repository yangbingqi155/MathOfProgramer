package main

import (
	"fmt"
	"math"
)

//Rand 随机数
const Rand = int32(37187)

//Dividend 除数
const Dividend = int32(7)

func main() {
	fmt.Println("实现一个简单的int32的数字的加密和解密函数\n")
	inputNumber := int32(625)

	encryptNumber, quotients := encrypt(inputNumber)
	fmt.Printf("明文数字 %d 加密后是密文数字是：%d\n", inputNumber, encryptNumber)
	fmt.Printf("明文数字 %d 加密后商分别是：", inputNumber)
	for j := 0; j < len(quotients); j++ {
		fmt.Printf(" , %d ", quotients[j])
	}
	fmt.Printf("\n")

	decryptNumber := decrypt(encryptNumber, quotients)
	fmt.Printf("密文数字 %d 解密成明文数字 %d \n", encryptNumber, decryptNumber)
}

//encrypt 加密函数
func encrypt(number int32) (int32, [3]int32) {
	tempNumber := number
	if tempNumber < 0 {
		number = -number
	}

	everyDigit, quotients := calculateDigitsAndQuotients(number)
	everyDigit = reverseNumber(everyDigit)

	encryptNumber := combineEveryDigitToNumber(everyDigit)

	if tempNumber < 0 {
		encryptNumber = -encryptNumber
	}

	return encryptNumber, quotients
}

//calculateDigitsAndQuotients 计算数字每位的值和每位的商(加密时使用，使用了随机数)
func calculateDigitsAndQuotients(number int32) (everyDigit [3]int32, quotients [3]int32) {
	k := number
	for i := len(everyDigit) - 1; i >= 0; i = i - 1 {
		digit := k / int32(math.Pow(10, float64(i)))
		quotients[len(everyDigit)-i-1] = (digit + Rand) / Dividend
		everyDigit[len(everyDigit)-i-1] = (digit + Rand) % Dividend
		k = k % int32(math.Pow(10, float64(i)))
	}
	return everyDigit, quotients
}

//getDigits 获取数字每位的值(只是获取每位的值)
func getDigits(number int32) (everyDigit [3]int32) {
	k := number
	for i := len(everyDigit) - 1; i >= 0; i = i - 1 {
		digit := k / int32(math.Pow(10, float64(i)))
		everyDigit[len(everyDigit)-i-1] = digit
		k = k % int32(math.Pow(10, float64(i)))
	}
	return everyDigit
}

//decrypt 解密函数
func decrypt(number int32, quotient [3]int32) (decryptNumber int32) {
	tempNumber := number
	if tempNumber < 0 {
		number = -number
	}
	//数字位反转
	everyDigit := getDigits(number)
	everyDigit = reverseNumber(everyDigit)

	for i := 0; i < len(everyDigit); i++ {
		everyDigit[i] = everyDigit[i] + quotient[i]*Dividend - Rand
	}
	decryptNumber = combineEveryDigitToNumber(everyDigit)
	if tempNumber < 0 {
		decryptNumber = -decryptNumber
	}
	return decryptNumber
}

//combineEveryDigitToNumber 放在数组里拆分的每一位组合成一个数字
func combineEveryDigitToNumber(everyDigit [3]int32) (number int32) {
	number = int32(0)
	for i := 0; i < len(everyDigit); i = i + 1 {
		number = number + everyDigit[i]*int32(math.Pow(10, float64(len(everyDigit)-i-1)))
	}
	return number
}

//reverseNumber 反转数字的每一位
func reverseNumber(numbers [3]int32) [3]int32 {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
