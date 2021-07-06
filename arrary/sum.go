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

func SumAll(numberToSum ...[]int) []int {
	// lengthOfNumbers := len(numberToSum)
	var sums []int
	for _, v := range numberToSum {
		sums = append(sums, Sum(v))

	}
	return sums
}

func SumTails(numberToSumTails ...[]int) []int {
	var sums []int
	for _, v := range numberToSumTails {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(v[1:]))

		}
	}
	return sums

}
