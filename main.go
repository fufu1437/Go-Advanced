package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// filterInts should return only the elements for which keep(v) is true,
// keeping their original order.
func filterInts(s []int, keep func(int) bool) []int {
	result := []int{}
	for _, v := range s {
		if keep(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	nums := []int{}
	for _, f := range strings.Fields(sc.Text()) {
		n, _ := strconv.Atoi(f)
		nums = append(nums, n)
	}

	// 只保留偶数
	evens := filterInts(nums, func(n int) bool { return n%2 == 0 })

	parts := make([]string, len(evens))
	for i, v := range evens {
		parts[i] = strconv.Itoa(v)
	}
	fmt.Println(strings.Join(parts, " "))
}
