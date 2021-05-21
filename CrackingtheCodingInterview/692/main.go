package main

import (
	"fmt"
	"sort"
)

//给一非空的单词列表，返回前 k 个出现次数最多的单词。

/**
输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
输出: ["i", "love"]
解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
    注意，按字母顺序 "i" 在 "love" 之前。
*/

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, value := range words {
		m[value]++
	}
	arr := make([]string, 0)
	for v, _ := range m {
		arr = append(arr, v)
	}
	sort.Slice(arr, func(i, j int) bool {
		if m[arr[i]] == m[arr[j]] {
			return arr[i] < arr[j]
		}
		return m[arr[i]] > m[arr[j]]
	})
	return arr[:k]
}
func main() {
	strs := []string{"i", "love", "leetcode", "i", "love", "coding"}
	frequent := topKFrequent(strs, 2)
	fmt.Println(frequent)
}
