package main

import "fmt"

/**
  题目：寻找最长不含有重复字符的字串
  思路：
     *定义一个map key存放字符串的每一个字符 value存对应下标
		在map中 key相同的话  后面设置的会替换掉前面的 不会重复

*/
func main() {
	fmt.Println(result("sddfgh")) //s:0 d:2 f:3 g:4 h:5
}

func result(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	// 遍历字符串
	for i, ch := range []byte(s) {
		//start 记录最长不重复字串的第一个字符的下表 如果在下一次遍历的时候 发现存在了当前遍历到的字符  并且这个字符的下标大于或者等于start
		//(即当前遍历到的字符和start位置对应的字符相同，要把start后移一位)
		if lastValue, temp := lastOccurred[ch]; temp && lastValue >= start {
			start = lastValue + 1
		}

		// 每一次都计算 maxLength的值
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i
	}

	fmt.Println(lastOccurred)
	return maxLength
}
