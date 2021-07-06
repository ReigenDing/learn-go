package main

func Sum(nums []int) int {
	sum := 0
	// for i := 0; i < len(nums); i++ {
	// 	sum += nums[i]
	// }
	// 使用range语法使得函数变得更加简介
	for _, v := range nums {
		sum += v

	}
	return sum
}

func SumAll(numberToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numberToSum)
	sums = make([]int, lengthOfNumbers)
	for i, v := range numberToSum {
		sums[i] = Sum(v)

	}
	return
}
