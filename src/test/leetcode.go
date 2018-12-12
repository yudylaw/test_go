package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func twoSum(nums []int, target int) []int {
	//O(N^2)
	size := len(nums)
	if size < 2 {
		return nil
	}

	for i, v1 := range nums {
		if i > size-2 {
			return nil
		}
		for j, v2 := range nums[i+1:] {
			if v1+v2 == target {
				return []int{i, j + i + 1}
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

func rand7() int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(7) + 1
	return num
}

func rand10() int {
	//使用平均2.3次rand7实现随机均匀的rand10，本质是一个概率问题
	a := rand7()
	b := rand7()
	if a > 4 && b < 4 {
		return rand10()
	} else {
		return (a+b)%10 + 1
	}
}

func isPalindrome(s string) bool {
	m := len(s) / 2
	//仅对比字母、数字，忽略大小写
	str := strings.ToLower(s)
	str2 := make([]int32, 0)
	for _, s1 := range str {
		if unicode.IsLetter(s1) || unicode.IsDigit(s1) {
			str2 = append(str2, s1)
		}
	}
	for i, s2 := range str2 {
		if i > m {
			break
		}
		j := len(str2) - 1 - i
		v1 := int32(str2[j])
		if s2 != v1 {
			return false
		}
	}
	return true
}

func intersection(nums1 []int, nums2 []int) []int {
	numsMap := make(map[int]int)
	for _, v1 := range nums1 {
		numsMap[v1] = 1
	}
	result := make([]int, 0)
	for _, v2 := range nums2 {
		if v, ok := numsMap[v2]; ok && v == 1 {
			numsMap[v2] = v + 1
			result = append(result, v2)
		}
	}
	return result
}

func main() {
	fmt.Println("hello leetcode.")
	//nums := []int{1, 5, 7, 2, 3, 4}
	//target := 9
	//arr := towSum_Nice(nums, target)
	//fmt.Printf("arr=%v", arr)

	//num := rand10()
	//fmt.Printf("num=%v", num)

	//str := "A man, a plan, a canal: Panama"
	////str := "Abba"
	//flag := isPalindrome(str)
	//fmt.Printf("flag=%v", flag)

	nums1 := []int{2, 1, 2, 3, 4, 5}
	nums2 := []int{5, 6, 1, 2, 3, 4, 5, 6}
	nums := intersection(nums1, nums2)
	fmt.Printf("nums=%v", nums)
}
