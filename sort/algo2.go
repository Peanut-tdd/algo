package main

import "fmt"

//搜索旋转排序数组
//输入: nums = [4,5,6,7,0,1,2], target = 0
//输出target 的数组下标: 4

//思路：二分查找，查询左右半边那边有序

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}

	k := arrKsearch(nums, 5)
	fmt.Println(k)

}

func arrKsearch(arr []int, target int) int {

	left, right := 0, len(arr)-1

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}

		if arr[left] <= arr[mid] {
			if arr[left] <= target && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else {
			if arr[mid] <= target && target < arr[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		}

	}
	return -1
}
