package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	arr := [1e4 + 10]bool{}
	for i := 1; i < 1e4+10; i++ {
		temp := i
		for temp > 0 {
			if temp%10 == 0 {
				arr[i] = false
				break
			}
			if i%(temp%10) != 0 {
				arr[i] = false
				break
			}
			temp /= 10
		}
		if temp == 0 {
			arr[i] = true
		}
	}
	ans := make([]int, 1e4)
	flag := 0
	for i := left; i <= right; i++ {
		if arr[i] {
			ans[flag] = i
			flag++
		}
	}
	return ans[:flag]
}
func main() {
	fmt.Println(selfDividingNumbers(1, 999))
}
