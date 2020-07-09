package main

import (
	"log"
	"os"
	"sort"

	"github.com/yuyamada/AtCoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, k := io.NextInt(), io.NextInt()
	a := io.NextInts(n)
	ans := solve(k, a)
	io.Println(ans)
}

// O(n log n)
func solve(k int, a []int) (ans int) {
	sort.Slice(a, func(i, j int) bool { return lib.Abs(a[i]) < lib.Abs(a[j]) })
	pNums, nNums := []int{}, []int{}
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] >= 0 {
			pNums = append(pNums, a[i])
		} else {
			nNums = append(nNums, a[i])
		}
	}

	ok := false
	if len(pNums) > 0 {
		if len(a) == k {
			ok = len(nNums)%2 == 0
		} else {
			ok = true
		}
	} else {
		ok = k%2 == 0
	}

	if !ok {
		ans = 1
		for i := 0; i < k; i++ {
			ans = lib.ModMul(ans, (a[i]+mod)%mod, mod)
		}
		return ans
	}

	ans = 1
	if k%2 == 1 {
		ans = lib.ModMul(ans, pNums[0], mod)
		pNums = pNums[1:]
	}
	nums := []int{}
	for len(pNums) >= 2 {
		nums = append(nums, pNums[0]*pNums[1])
		pNums = pNums[2:]
	}
	for len(nNums) >= 2 {
		nums = append(nums, nNums[0]*nNums[1])
		nNums = nNums[2:]
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	for i := 0; i < k/2; i++ {
		ans = lib.ModMul(ans, nums[i], mod)
	}
	return ans
}

const mod = 1e9 + 7

var logger = log.New(os.Stderr, "", 0)
