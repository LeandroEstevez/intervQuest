package main

import "fmt"

func main() {

}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var i int = 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	elements := make(map[int]int)
	for i, v := range nums {
		_, ok := elements[v]
		if ok {
			return true
		}
		elements[v] = i
	}
	return false
}

func singleNumber(nums []int) int {
	elements := make(map[int]int)
	for _, v := range nums {
		_, ok := elements[v]
		if ok {
			elements[v] = elements[v] + 1
		} else {
			elements[v] = 1
		}
	}
	for _, v := range nums {
		if elements[v] == 1 {
			return v
		}
	}

	return nums[0]
}

func intersect(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]int)
	nums2Map := make(map[int]int)

	for _, v := range nums1 {
		_, ok := nums1Map[v]
		if ok {
			nums1Map[v] = nums1Map[v] + 1
		} else {
			nums1Map[v] = 1
		}
	}

	for _, v := range nums2 {
		_, ok := nums2Map[v]
		if ok {
			nums2Map[v] = nums2Map[v] + 1
		} else {
			nums2Map[v] = 1
		}
	}

	fmt.Println(nums1Map)
	fmt.Println(nums2Map)

	var comm []int

	if len(nums1) < len(nums2) {

		for k, v := range nums1Map {

			i := 0
			n2, ok := nums2Map[k]

			if ok {
				if n2 < v {
					for ;i < n2; i++ {
						comm = append(comm, k)
						fmt.Println("Doing n")
					}
				} else {
					for ;i < v; i++ {
						comm = append(comm, k)
						fmt.Println("Doing m")
					}
				}
			}

		}

	} else {

		for k, v := range nums2Map {

			i := 0
			n1, ok := nums1Map[k]

			if ok {
				if n1 < v {
					for ;i < n1; i++ {
						comm = append(comm, k)
						fmt.Println("Doing n")
					}
				} else {
					for ;i < v; i++ {
						comm = append(comm, k)
						fmt.Println("Doing m")
					}
				}
			}

		}

	}

	return comm
}
