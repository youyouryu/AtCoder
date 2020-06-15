package lib

import (
	"strconv"
)

// Ints2Strs converts int slice to string slice
func Ints2Strs(slice []int) []string {
	strs := []string{}
	for _, v := range slice {
		strs = append(strs, strconv.Itoa(v))
	}
	return strs
}
