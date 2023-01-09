package functions

import (
	"StudyDemo/CurriculumDesign/Sort/utils"
	"time"
)

func ShellSort(array [80000]int) (ConstTime int64) {
	start := time.Now().UnixMilli()
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
	end := time.Now().UnixMilli()
	sli := array[:]
	utils.StoreFiles(sli, "ShellSort")
	return end - start
}
