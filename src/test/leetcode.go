package main

import (
	"fmt"
	"math"
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

func isSubsequence(s string, t string) bool {
	//双指针法
	lt := len(t)
	ls := len(s)
	if ls > lt {
		return false
	}
	i := 0
	for _, v := range t {
		if i < ls && int32(s[i]) == v {
			i++
		}
	}
	return i == ls
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 中序遍历
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	list := make([]int, 0)
	//递归实现
	//先中序遍历左子树，然后再访问根结点，最后再中序遍历右子树
	//非递归实现：基于栈
	leftList := inorderTraversal(root.Left)
	if leftList != nil {
		list = append(list, leftList...)
	}
	list = append(list, root.Val)
	rightList := inorderTraversal(root.Right)
	if rightList != nil {
		list = append(list, rightList...)
	}
	return list
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	//翻转左右子树，再替换
	leftTree := invertTree(root.Left)
	rightTree := invertTree(root.Right)

	root.Left = rightTree
	root.Right = leftTree
	return root
}

//层次遍历，广度优先搜索
func levelOrderSimple(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	//使用数组实现的队列
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	//返回无层次的一维数组，每次添加一个节点
	values := make([]int, 0)
	for len(queue) > 0 {
		node := queue[0]
		values = append(values, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		//弹出队头
		queue = queue[1:]
	}
	return values
}

//层次遍历，广度优先搜索
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	list := make([][]int, 0)
	//使用数组实现的队列
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		//父节点长度
		count := len(queue)
		values := make([]int, 0)
		//遍历所有子节点
		for i := 0; i < count; i++ {
			node := queue[i]
			values = append(values, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		//弹出上一层所有父节点
		queue = queue[count:]
		list = append(list, values)
	}
	return list
}

//插入二叉搜索树
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := root
	parent := root
	newNode := &TreeNode{Val: val}
	for node != nil {
		parent = node
		if val > node.Val {
			//right tree
			node = node.Right
			if node == nil {
				parent.Right = newNode
				break
			}
		} else {
			//left tree
			node = node.Left
			if node == nil {
				parent.Left = newNode
				break
			}
		}
	}
	return root
}

func searchBST(root *TreeNode, val int) *TreeNode {
	node := root
	parent := root
	flag := false
	for node != nil {
		parent = node
		if val > node.Val {
			//right tree
			node = node.Right
		} else if val < node.Val {
			//left tree
			node = node.Left
		} else {
			flag = true
			break
		}
	}
	if flag {
		return parent
	} else {
		return nil
	}
}

func isBalanced(root *TreeNode) bool {
	//平衡二叉树条件:
	//1:左右子树的高度差不大于1
	//2:并且左右子树都是平衡二叉树
	if root == nil {
		return true
	}
	leftHigh := maxTreeDepth(root.Left)
	rightHigh := maxTreeDepth(root.Right)
	diff := leftHigh - rightHigh
	//左右子树的高度差不大于1
	if math.Abs(float64(diff)) > 1 {
		return false
	} else {
		//并且左右子树都是平衡二叉树
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
}

func maxTreeDepth(root *TreeNode) int {
	//需要算上父节点高度 +1
	if root != nil {
		lhigh := maxTreeDepth(root.Left)
		rhigh := maxTreeDepth(root.Right)
		if lhigh > rhigh {
			return lhigh + 1
		} else {
			return rhigh + 1
		}
	} else {
		return 0
	}
}

func reverseString(s string) string {
	str := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		str = append(str, s[i])
	}
	return string(str[:])
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

	//nums1 := []int{2, 1, 2, 3, 4, 5}
	//nums2 := []int{5, 6, 1, 2, 3, 4, 5, 6}
	//nums := intersection(nums1, nums2)
	//fmt.Printf("nums=%v", nums)

	//s := "abc"
	//t := "abbc"
	//flag := isSubsequence(s, t)
	//fmt.Printf("flag=%v", flag)

	//二叉树 binary tree
	//n8 := &TreeNode{Val: 8}
	//n7 := &TreeNode{Val: 7, Right: n8}
	//n6 := &TreeNode{Val: 6}
	//n5 := &TreeNode{Val: 5, Left: n7}
	//n4 := &TreeNode{Val: 4}
	//n3 := &TreeNode{Val: 3, Right: n6}
	//n2 := &TreeNode{Val: 2, Left: n4, Right: n5}
	//root := &TreeNode{Val: 1, Left: n2, Right: n3}
	//
	//list := levelOrder(root)
	//fmt.Printf("list=%v", list)
	//newTree := invertTree(root)
	//list := inorderTraversal(newTree)
	//for _, val := range list {
	//	fmt.Printf("%d,", val)
	//}

	//二叉搜索树 binary search tree(BST)
	n7 := &TreeNode{Val: 7}
	n3 := &TreeNode{Val: 3}
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2, Left: n1, Right: n3}
	root := &TreeNode{Val: 4, Left: n2, Right: n7}

	flag := isBalanced(root)
	fmt.Printf("flag=%v", flag)

	//newTree := insertIntoBST(root, 5)
	//newTree := searchBST(root, 2)
	//list := inorderTraversal(newTree)
	//for _, val := range list {
	//	fmt.Printf("%d,", val)
	//}

	//str := string("ab,cdff :f")
	//newStr := reverseString(str)
	//fmt.Printf("%v", newStr)
}
