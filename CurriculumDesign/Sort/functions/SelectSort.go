package functions

import (
	"StudyDemo/CurriculumDesign/Sort/utils"
	"time"
)

// SelectSort 假定第一个就是最小的，找出最小值，最小值和无序序列的第一个值位置互换，不停地把最小地元素放在整个无序序列的最前面
func SelectSort(array [80000]int) (ConstTime int64) {
	start := time.Now().UnixMilli()
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
	end := time.Now().UnixMilli()
	sli := array[:]
	utils.StoreFiles(sli, "SelectSort")
	return end - start
}
