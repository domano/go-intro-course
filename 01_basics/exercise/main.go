//Input 4

// 1   2   3   4
//  2   4   6   8
//  3   6   9  12
//  4   8  12  16

package main

import "fmt"

func main() {
	lines := multiplcationTable(4)
	fmt.Printf("%v", lines)
}

func multiplcationTable(n int) [][]int {
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	

	lines := [][]int{}
	lines.typeof
	for _, num := range nums {
		line := []int{}
		for _, k := range nums {
			line = append(line, num*k)
		}
		lines = append(lines, line)
	}
	return lines
}
