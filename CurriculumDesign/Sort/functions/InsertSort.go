package functions

import (
	"StudyDemo/CurriculumDesign/Sort/utils"
	"time"
)

func InsertSort(array [80000]int) (ConstTime int64) {
	start := time.Now().UnixMilli()
	for i := range array {
		preIndex := i - 1
		current := array[i]
		for preIndex >= 0 && array[preIndex] > current {
			array[preIndex+1] = array[preIndex]
			preIndex--
		}
		array[preIndex+1] = current
	}
	end := time.Now().UnixMilli()
	sli := array[:]
	utils.StoreFiles(sli, "InsertSort")
	return end - start
}
