package functions

import (
	"StudyDemo/CurriculumDesign/Sort/utils"
	"time"
)

// 插入排序
// 从小到大
func InsertSortSmallestToLargest(array [80000]int) (ConstTime int64) {
	start := time.Now().UnixNano()
	for i := range array {
		preIndex := i - 1
		current := array[i]
		for preIndex >= 0 && array[preIndex] > current {
			array[preIndex+1] = array[preIndex]
			preIndex--
		}
		array[preIndex+1] = current
	}
	end := time.Now().UnixNano()
	sli := array[:]
	utils.StoreFiles(sli, "InsertSortSmallestToLargest")
	return end - start
}

// 从大到小
func InsertSortLargestToSmallest(array [80000]int) (ConstTime int64) {
	start := time.Now().UnixNano()
	for i := range array {
		preIndex := i - 1
		current := array[i]
		for preIndex >= 0 && array[preIndex] < current {
			array[preIndex+1] = array[preIndex]
			preIndex--
		}
		array[preIndex+1] = current
	}
	end := time.Now().UnixNano()
	sli := array[:]
	utils.StoreFiles(sli, "InsertSortLargestToSmallest")
	return end - start
}
