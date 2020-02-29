package main

func main() {

}

/*
24 ms  38.46%
5.9 MB  100.00%
*/
func missingNumber(nums []int) int {
	lens := len(nums)
	arr := make([]bool, lens+1)
	for _, v := range nums {
		arr[v] = true
	}
	for i, v := range arr {
		if v == false {
			return i
		}
	}
	return 0
}
