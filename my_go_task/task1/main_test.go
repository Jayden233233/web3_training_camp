package main

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums := []int{2, 2, 1}
	fmt.Printf("singleNumber(nums): %v\n", singleNumber(nums))
}

func TestIsPalindrome(t *testing.T) {
	x := 121
	fmt.Printf("isPalindrome(x): %v\n", isPalindrome(x))

}

// 136. 只出现一次的数字
// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
func singleNumber(nums []int) int {

	res := 0
	for _, x := range nums {
		res ^= x
	}
	return res
}

// 9. 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 例如，121 是回文，而 123 不是。

// 收获 ：遍历字符串时要区分asci字符串和unicode字符串
func isPalindrome(x int) bool {
	num_str := strconv.Itoa(x)
	chars_slice := []rune(num_str)

	i := 0
	j := len(chars_slice) - 1
	for i < j {
		if chars_slice[i] != chars_slice[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 20.有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。

// 收获，整理了切片的最佳实践和字符串的最佳实践
func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if (v == ')' && top == '(') || (v == ']' && top == '[') || (v == '}' && top == '{') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return true
}

// 14. 最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
func longestCommonPrefix(strs []string) string {
	if len(strs[0]) == 0 {

		return ""
	}
	for i := 0; i < len(strs[0]); i++ {

		curChar := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != curChar {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// 66. 加一
func plusOne(digits []int) []int {
	var jinwei int = 1
	for i := len(digits) - 1; i >= 0; i-- {
		var temp int = digits[i] + jinwei
		yu := temp % 10
		jinwei = temp / 10
		digits[i] = yu
	}
	if jinwei == 1 {

		return append([]int{1}, digits...)
	} else {

		return digits
	}
}

// 26. 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	var k = 0
	for i := 1; i < len(nums); i++ {
		if nums[k] == nums[i] {
			continue
		} else {
			nums[k+1] = nums[i]
			k++
		}
	}
	return k + 1
}

// 56. 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 按区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}
	for _, current := range intervals[1:] {
		// 获取merged中最后一个区间
		last := merged[len(merged)-1]

		// 如果当前区间与最后一个区间重叠，则合并
		if current[0] <= last[1] {
			// 更新最后一个区间的结束位置为两者的最大值
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 不重叠，添加新区间
			merged = append(merged, current)
		}
	}

	return merged
}

// 1. 两数之和
func twoSum(nums []int, target int) []int {
	nummap := make(map[int]int, 0)
	for i, _ := range nums {

		index, exists := nummap[target-nums[i]]
		if exists {
			return []int{index, i}
		} else {
			nummap[nums[i]] = i
		}
	}
	return []int{}
}
