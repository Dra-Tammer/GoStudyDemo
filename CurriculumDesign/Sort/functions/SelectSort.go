package functions

import (
	"StudyDemo/CurriculumDesign/Sort/utils"
	"time"
)

// 选择排序
// 从小到大
func SelectSortSmallestToLargest(array [80000]int) (ConstTime int64) {
	start := time.Now().UnixNano()
	var k, temp int
	for i := 0; i < len(array)-1; i++ {
		k = i
		for j := i + 1; j < len(array); j++ {
			if array[k] > array[j] {
				k = j //寻找最小的数，将最小的数的搜引保存
			}
		}
		temp = array[k]
		array[k] = array[i]
		array[i] = temp
	}
	end := time.Now().UnixNano()
	sli := array[:]
	utils.StoreFiles(sli, "SelectSortSmallestToLargest")
	return end - start
}

// 从大到小
func SelectSortLargestToSmallest(array [80000]int) (ConstTime int64) {
	start := time.Now().UnixNano()
	var k, temp int
	for i := 0; i < len(array)-1; i++ {
		k = i
		for j := i + 1; j < len(array); j++ {
			if array[k] < array[j] {
				k = j
			}
		}
		temp = array[k]
		array[k] = array[i]
		array[i] = temp
	}
	end := time.Now().UnixNano()
	sli := array[:]
	utils.StoreFiles(sli, "SelectSortLargestToSmallest")
	return end - start
}
