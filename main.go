package main

import (
	"fmt"
)

func main() {
	reverse(1)
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

func plusOne(digits []int) []int {
	if digits[(len(digits) - 1)] != 9 {
		digits[(len(digits) - 1)] = digits[(len(digits) - 1)] + 1
	} else {
		for i := len(digits) - 1; i >= 0; i-- {
			if i == 0 && digits[i] == 9 {
				digits = append(digits, 0)
				digits[i] = 1
			} else {
				if digits[i] != 9 {
					digits[i] = digits[i] + 1
					break
				} else {
					digits[i] = 0
				}
			}
		}
	}
	return digits
}

func moveZeroes(nums []int)  {
	var j int
	var z int
	if len(nums) != 1 {
		for i := 0; i < len(nums); i++ {
			j = i + 1
			if i == len(nums) - 1 {
				break
			}
			if nums[i] == 0 {
				z = i
				for nums[j] != 0 {
					nums[z] = nums[j]
					nums[j] = 0
					z = j
					j++
				}
			}
		}
	}
}

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int][]int)

	var result []int

	for i, v := range nums {
		_, ok := numsMap[v]
		if !ok {
			numsMap[v] = []int{i}
		} else {
			numsMap[v] = append(numsMap[v], i)
		}
	}

	for i, v := range nums {
		diff := target - v
		j, ok := numsMap[diff]
		if ok {
			if i == j[0] {
				if len(j) > 1 {
					result = append(result, i, j[1])
					break
				}
				continue
			}
			result = append(result, i, j[0])
			break
		}
	}

	return result
}

func reverseString(s []byte)  {
	i := 0
	j := len(s) - 1
	for i < j {
		temp := s[j]
		s[j] = s[i]
		s[i] = temp
		i++
		j--
	}
}

func firstUniqChar(s string) int {
	sMap := make(map[rune][]int)
	for i, v := range s {
		sMap[v] = append(sMap[v], i)
	}
	for _, v := range s {
		if len(sMap[v]) == 1 {
			return sMap[v][0]
		}
	}
	return -1
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int)
	for _, v := range s {
		sMap[v] = sMap[v] + 1
	}
	fmt.Println(sMap)
	tMap := make(map[rune]int)
	for _, v := range t {
		tMap[v] = tMap[v] + 1
	}
	fmt.Println(tMap)
	count := 0
	for k, v := range sMap {
		n, ok := tMap[k]
		if ok && v == n {
			count = count + n
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
