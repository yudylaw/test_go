package main

import "fmt"

func twoSum(nums []int, target int) []int {
	//O(N^2)
	size := len(nums)
	if size < 2 {
		return nil
	}

	for i, v1 := range nums {
		if i > size - 2 {
			return nil
		}
		for j, v2 := range nums[i+1:] {
			if v1 + v2 == target {
				return []int{i,j+i+1}
			}
		}
	}
	return nil
}

func towSum_Nice(nums []int, target int) []int {
	//O(N)
	size := len(nums)
	if size < 2 {
		return nil
	}

	remain := make(map[int]int)

	for i, v1 := range nums {
		s := target - v1 //求差
		if v, ok := remain[s]; ok {
			return []int{v, i}
		}
		remain[v1] = i
	}

	return nil
}

func main() {
	nums := []int{1,5,7,2,3,4}
	target := 9
	arr := towSum_Nice(nums, target)
	fmt.Printf("arr=%v", arr)
}
