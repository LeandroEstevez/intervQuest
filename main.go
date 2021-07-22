package main

import "fmt"

func main() {
	numbers := []int{1,1,2}
	n := removeDuplicates(numbers)
	fmt.Println(n)
}

func removeDuplicates(nums []int) int {
	c := 0
	var i int = 0
	var n int = nums[i]
	for j := 1; j < len(nums); j++ {

		if n == nums[j] {
			fmt.Println("Increasing count")
			c++
		} else {
			i++
			nums[i] = nums[j]
			fmt.Println(nums)
			n = nums[j]
			fmt.Println("n =", n)
		}
	}

	return c
}
