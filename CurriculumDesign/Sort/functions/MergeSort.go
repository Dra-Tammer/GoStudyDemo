package functions

import (
	"StudyDemo/CurriculumDesign/Sort/utils"
	"time"
)

// MergeSort 先拆分，然后一层一层向上合并，保证左边的序列和右边的序列都是有序的，最后一层一层地向上合并
func MergeSort(arr []int) (ConstTime int64) {
	start := time.Now().UnixMilli()
	result := mergeSort(arr)
	end := time.Now().UnixMilli()
	utils.StoreFiles(result, "MergeSort")
	return end - start
}

func mergeSort(arr []int) []int {
	length := len(arr)
	//如果只有一个元素，那么就不需要继续划分
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(mergeSort(left), mergeSort(right)) //把左半区和右半区合并起来，括号里面是对区域的不断拆分
}

func merge(left []int, right []int) []int {
	var result []int
	//左边和右边都有元素的情况
	for len(left) != 0 && len(right) != 0 { //左边区域的第一个元素和右边区域的第一个元素比较，谁小谁往result数组里放
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:] //更新左边区域的内容
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	//只有左边有元素，最后留着的数是有序的，直接往里面添加即可
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	//只有右边有元素
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}
