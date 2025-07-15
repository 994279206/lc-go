package leetcode

import "fmt"

//统计n以内的素数个数
//暴力破解，便利所有数字，找出只能够被自身和1整除的数据。0，1除外
//埃筛法

func func001(n int) int {
	isPrime := map[int]bool{}
	for i := 2; i <= n; i++ {
		isPrime[i] = false
	}
	ans := 0
	for i := 2; i < n; i++ {
		if !isPrime[i] {
			ans += 1
			for j := i * i; j < n; j += i {
				isPrime[j] = true
			}
		}
	}
	fmt.Println(ans)
	return ans
}
