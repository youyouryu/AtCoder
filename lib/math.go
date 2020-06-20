package lib

// Abs retuns absolutevalue of input
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Max returns max value of inputs
func Max(a ...int) (max int) {
	for i, ai := range a {
		if i == 0 || ai > max {
			max = ai
		}
	}
	return max
}

// Min returns minimum value of inputs
func Min(a ...int) (min int) {
	for i, ai := range a {
		if i == 0 || ai < min {
			min = ai
		}
	}
	return min
}

// ModPow returns base^n with mod
func ModPow(base, n, mod int) int {
	ans := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			ans = ans * base % mod
		}
		base = base * base % mod
	}
	return ans
}
