package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	// TODO: give squares a capacity of n up front instead of starting empty,
	// then read n values and append the square of each one.
	squares := make([]int, 0, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		squares = append(squares, v*v)
	}

	// Printing is already wired up for you.
	parts := make([]string, len(squares))
	for i, s := range squares {
		parts[i] = strconv.Itoa(s)
	}
	fmt.Println(strings.Join(parts, " "))
}
