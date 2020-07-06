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

// Gcd returns greatest common divisor of a and b
func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// ModMul returns a*b with mod
func ModMul(a, b, mod int) int {
	a, b = a%mod, b%mod
	if b == 0 {
		return 0
	}
	if a*b/b == a {
		return a * b % mod
	}
	panic("overflow")
}

// ModPow returns base^n with mod
func ModPow(base, n, mod int) int {
	ans := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			ans = ModMul(ans, base, mod)
		}
		base = ModMul(base, base, mod)
	}
	return ans
}

// ModInv returns inverse value of a with mod
func ModInv(a, mod int) int {
	return ModPow(a, mod-2, mod)
}

// ModFactorial stores factorials with mod
type ModFactorial struct {
	mod    int
	values []int
	invs   []int
}

// NewModFactorial returns new instance of ModFactorial
func NewModFactorial(mod int) *ModFactorial {
	return &ModFactorial{
		mod:    mod,
		values: []int{1},
		invs:   []int{1},
	}
}

// Calc returns n!
func (m *ModFactorial) Calc(n int) int {
	l := len(m.values)
	if n < l {
		return m.values[n]
	}
	tmp := make([]int, n+1-l)
	m.values = append(m.values, tmp...)
	for i := l; i <= n; i++ {
		m.values[i] = i * m.values[i-1] % m.mod
	}
	return m.values[n]
}

// CalcInv returns inverse value of n!
func (m *ModFactorial) CalcInv(n int) int {
	l := len(m.invs)
	if n < l {
		return m.invs[n]
	}
	tmp := make([]int, n+1-l)
	m.invs = append(m.invs, tmp...)
	m.invs[n] = ModInv(m.Calc(n), m.mod)
	for i := n - 1; i >= l; i-- {
		m.invs[i] = m.invs[i+1] * (i + 1) % m.mod
	}
	return m.invs[n]

}

// ModCombination returns the number of combinations with mod
type ModCombination struct {
	mod       int
	factorial *ModFactorial
}

// NewModCombination returns new instance of ModCombination
func NewModCombination(mod int) *ModCombination {
	return &ModCombination{
		mod:       mod,
		factorial: NewModFactorial(mod),
	}
}

// Calc returns nCk
func (m *ModCombination) Calc(n, k int) int {
	if n < k || n < 0 || k < 0 {
		return 0
	}
	fn := m.factorial.Calc(n)
	ifk := m.factorial.CalcInv(k)
	ifnk := m.factorial.CalcInv(n - k)
	ans := ModMul(fn, ifk, m.mod)
	ans = ModMul(ans, ifnk, m.mod)
	return ans
}
