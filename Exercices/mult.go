package main

import "fmt"

func main() {
	// clasic for
	fmt.Println(ltabmult(4,1))

	fmt.Println(tabmult(4))
}

func ltabmult( n int, par int) [] int{
	nums := []int{}
	for i:=1; i<= n; i++ {
        nums = append(nums, i*par)
	}
	return nums
}

func tabmult(n int) [][] int {
	lines := [][]int{}

	for i:=1; i<= n; i++ {
		lines = append (lines,ltabmult(n, i))
	}

	return lines

}
