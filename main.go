package main

import "fmt"

func main() {
	//var N int
	//fmt.Scanf("%d", &N)

	//fmt.Println(s)

	ht := make(map[int][]int)

	var n int
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		ht[x] = append(ht[x], y)
	}

	print("xxx")

	//arr := make([]int, N)
	//sum := 0
	//for i := 0; i < N; i++ {
	//	fmt.Scanf("%d", &arr[i])
	//	sum += arr[i]
	//}
	//
	//// sum = 红 + 蓝 + 白 => 红 = 蓝 = （sum - 白）/2
	//// 动态规划  dp[i][j] 表示 前 i 个数中 红 -蓝 = j 的方案数，范围 -S S  由于不支持 统一 +S
	//// 初始化 dp[0][sum] = 1
	//
	//dp := make([][]int, N+1)
	//for i := 0; i <= N; i++ {
	//	dp[i] = make([]int, 2*sum+1)
	//}
	//
	//dp[0][sum] = 1 // 表示从 前 0 个钟选择 红-蓝 = 0 的方案数 为 1
	//// 遍历物品
	//for i := 1; i <= N; i++ {
	//	for j := 0; j <= 2*sum; j++ {
	//		// 如果当前不涂色 + 涂红 + 涂蓝
	//		dp[i][j] = dp[i-1][j]
	//		if j-arr[i-1] >= 0 {
	//			dp[i][j] += dp[i-1][j-arr[i-1]]
	//		}
	//		if j+arr[i-1] <= 2*sum {
	//			dp[i][j] += dp[i-1][j+arr[i-1]]
	//		}
	//	}
	//}
	//// -1 是因为全不涂
	//fmt.Println(dp[N][sum] - 1)
}
