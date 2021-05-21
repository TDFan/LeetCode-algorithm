package main

import "fmt"

//两个整数的 汉明距离 指的是这两个数字的二进制数对应位不同的数量。
/**
输入: 4, 14, 2
输出: 6
HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.

*/
func totalHammingDistance(nums []int) int {
	arr := [32]int{}
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		flag := 0
		for temp > 0 {
			if temp&1 == 1 {
				arr[flag]++
			}
			flag++
			temp >>= 1
		}
	}
	ans := 0
	for i := 0; i < 32; i++ {
		ans += (len(nums) - arr[i]) * arr[i]
	}
	return ans
}
func main() {
	arr := []int{4, 14, 2}
	fmt.Println(totalHammingDistance(arr))
}
