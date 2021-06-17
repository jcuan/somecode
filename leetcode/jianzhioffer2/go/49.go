package jzoffer

func min(a, b, c int) int {
	minval := a
	if b < minval {
		minval = b
	}
	if c < minval {
		minval = c
	}
	return minval
}

func nthUglyNumber(n int) int {
	p2, p3, p5 := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		p2val := dp[p2] * 2
		p3val := dp[p3] * 3
		p5val := dp[p5] * 5
		minval := min(p2val, p3val, p5val)
		if p2val == minval {
			p2++
		}
		if p3val == minval {
			p3++
		}
		if p5val == minval {
			p5++
		}
		dp[i] = minval
	}
	return dp[n-1]
}
