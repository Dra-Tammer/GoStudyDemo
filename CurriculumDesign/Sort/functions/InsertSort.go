package functions

import (
	"StudyDemo/CurriculumDesign/Sort/utils"
	"time"
)

// InsertSort 将小的元素不停地往前插入，比他大的元素右移
func InsertSort(array [80000]int) (ConstTime int64) {
	start := time.Now().UnixMilli()
	for i := range array {
		preIndex := i - 1 //第一个元素本身是有序的，对其不进行插入排序
		current := array[i]
		for preIndex >= 0 && array[preIndex] > current {
			array[preIndex+1] = array[preIndex] //比他的元素右移
			preIndex--
		}
		//发现此时preindex不满足比current小的条件，把current插入到他的后面
		array[preIndex+1] = current
	}
	end := time.Now().UnixMilli()
	sli := array[:]
	utils.StoreFiles(sli, "InsertSort")
	return end - start
}
