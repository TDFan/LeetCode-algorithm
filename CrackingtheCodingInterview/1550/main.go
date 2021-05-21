package main

import "fmt"

//给你一个整数数组 arr，请你判断数组中是否存在连续三个元素都是奇数的情况：如果存在，请返回 true ；否则，返回 false 。
/**
输入：arr = [2,6,4,1]
输出：false
解释：不存在连续三个元素都是奇数的情况。
*/
func threeConsecutiveOdds(arr []int) bool {
	begin := 0
	end := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]&1 == 1 {
			end = i
		} else {
			begin = i
		}
		if end-begin >= 3 {
			return true
		}
	}
	return false
}
func main() {
	arr := []int{2, 6, 4, 1}
	fmt.Println(threeConsecutiveOdds(arr))
}
