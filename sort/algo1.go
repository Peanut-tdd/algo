package main

import "fmt"

// 有序数组，查找target元素出现次数
//统计数字出现次数
//题目描述
//给定一个有序数组和一个目标数字 K，统计 K 在数组中出现的次数
//示例：
//输入：lists = [1,2,3,3,3,3,3,4], K= 3
//输出：5
//输入：lists = [1,3,3], K = 4
//输出：0
//其中数组长度 0 <= length <= 10 ^ 5

//设计思路：二分查找，查target最左最右的偏移量:right-left+1

func main() {
	arr := []int{1, 2, 3, 3, 3, 3, 3, 4}
	fmt.Println(targetSearchCount(arr, 3)) //5
	fmt.Println(targetSearchCount(arr, 4)) //1
	fmt.Println(targetSearchCount(arr, 5)) //0
}

func targetSearchCount(slice []int, target int) int {

	if len(slice) == 0 {
		return 0
	}

	//左半边
	leftIndex := func(arr []int, num int) int {
		left, right := 0, len(arr)-1
		res := -1
		for left <= right {
			mid := left + (right-left)/2

			if arr[mid] == num {
				res = mid
				right = mid - 1
			} else if arr[mid] < num { //left 左移
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return res

	}(slice, target)

	rightIndex := func(arr []int, num int) int {
		left, right := 0, len(arr)-1
		res := -1
		for left <= right {
			mid := left + (right-left)/2
			if arr[mid] == num {
				res = mid
				left = mid + 1
			} else if arr[mid] < num {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		return res
	}(slice, target)

	if leftIndex == -1 {
		return 0
	}

	return rightIndex - leftIndex + 1
}
