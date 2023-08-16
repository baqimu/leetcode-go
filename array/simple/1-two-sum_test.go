package simple

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Printf("两数只和结果：%v", twoSum(nums, target))
	fmt.Printf("使用hash表，两数只和结果：%v", twoSumhash(nums, target))
}

// 方法1： 暴力循环
func twoSum(nums []int, target int) []int {
	// 初始化 一个固定长度数组
	result := make([]int, 2)

	for first, value1 := range nums {
		for second, value2 := range nums {

			if value1+value2 == target && first != second {
				result[0] = first
				result[1] = second
				// break 只能中止一次for，switch 循环
				break
			}
		}
		// 二次中止 break
		if result[1] != 0 {
			break
		}
	}

	return result
}

// 方法二： 使用hash表,key 存储数组值；值存储数组下标。 通过 target-x 进行判定另一个 是否存在
func twoSumhash(nums []int, target int) []int {

	var hashtable map[int]int
	hashtable = make(map[int]int)
	for i, v := range nums {
		if key, value := hashtable[target-v]; value {
			return []int{key, i}
		}
		hashtable[v] = i
	}
	return nil
}
