package main

import "fmt"

func waysToStep(n int) int {
	a, b := 0, 0
	c := 1
	for n > 0 {
		a, b, c = b, c, (a+b+c)%1000000007
		n--
	}
	return c
}
func main() {
	ans := waysToStep(5)
	fmt.Println(ans)
}
