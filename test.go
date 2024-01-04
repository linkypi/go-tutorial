/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2023-07-16 13:35:14
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2024-01-04 17:59:44
 * @FilePath: /undefined/Users/leo/Documents/go.projects/test/test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"strings"
)

func main34() {

	intBreak(10)
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lisx(arr)
	fmt.Println("logest increase len: ", res)

	strs := []string{"flow", "flower", "flight"}
	prefix := longestCommonPrefix(strs)
	fmt.Println("logest common prefix: ", prefix)
}

// 将一个整数拆分为多个正整数之和,并使得这些正整数的乘积最大
func intBreak(n int) int {
	dp := []int{}
	dp = make([]int, n+1)

	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			// dp[i] 表示正整数 i 拆分为两个正整数 j 及 (i-j) 后可以获取的最大乘积
			// 1. 若 j 可以再拆分, 则乘积为 dp[j]* (i-j)
			// 2. 若 j 不可再拆分, 则 j*(i-j)
			tmp := max(dp[j]*(i-j), j*(i-j))
			dp[i] = max(dp[i], tmp)
		}
	}
	return dp[n]
}

func lisx(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]

	for _, k := range strs {
		index := strings.Index(k, prefix)
		for index != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
