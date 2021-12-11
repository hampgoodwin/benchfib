package fib

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibMemo(n int, memo []int) int {
	if n < 2 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = fibMemo(n-2, memo) + fibMemo(n-1, memo)

	return memo[n]
}

func fibBottomUp(n int) int {
	if n < 2 {
		return n
	}
	bottom_up := make([]int, n+1)
	bottom_up[1] = 1
	bottom_up[2] = 1
	for i := 3; i < n; i++ {
		bottom_up[i] = bottom_up[i-1] + bottom_up[i-2]
	}
	return bottom_up[n]
}
