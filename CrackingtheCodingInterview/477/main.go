package main

import "fmt"

//两个整数的 汉明距离 指的是这两个数字的二进制数对应位不同的数量。
/**
输入: 4, 14, 2
输出: 6
HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.

*/
func HammingDistance(a, b int) int {
	temp := a ^ b
	sum := 0
	for {
		if temp&1 == 1 {
			sum++
		}
		temp >>= 1
		if temp == 0 {
			break
		}
	}
	return sum
}
func totalHammingDistance(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			distance := HammingDistance(nums[i], nums[j])
			ans += distance
		}
	}
	return ans
}
func main() {
	arr := []int{4, 14, 2}
	fmt.Println(totalHammingDistance(arr))
}
