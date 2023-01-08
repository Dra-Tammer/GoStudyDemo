package functions

import (
	"StudyDemo/CurriculumDesign/Sort/utils"
	"time"
)

func ShellSortSmallestToLargest(array [80000]int) (ConstTime int64) {
	start := time.Now().UnixNano()
	length := len(array)
	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := array[i]
			j := i - gap
			for j >= 0 && array[j] > temp {
				array[j+gap] = array[j]
				j -= gap
			}
			array[j+gap] = temp
		}
		gap = gap / 3
	}
	end := time.Now().UnixNano()
	sli := array[:]
	utils.StoreFiles(sli, "ShellSortSmallestToLargest")
	return end - start
}

func ShellSortLargestToSmallest(array [80000]int) (ConstTime int64) {
	start := time.Now().UnixNano()
	length := len(array)
	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := array[i]
			j := i - gap
			for j >= 0 && array[j] < temp {
				array[j+gap] = array[j]
				j -= gap
			}
			array[j+gap] = temp
		}
		gap = gap / 3
	}
	end := time.Now().UnixNano()
	sli := array[:]
	utils.StoreFiles(sli, "ShellSortLargestToSmallest")
	return end - start
}
