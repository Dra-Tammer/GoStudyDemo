package functions

import (
	"StudyDemo/CurriculumDesign/Sort/utils"
	"time"
)

// 冒泡排序
// 从小到大
func BubblingSortSmallestToLargest(InitedArray [80000]int) (ConstTime int64) { //从小到大排序
	array := InitedArray
	start := time.Now().UnixNano() //获取当前的时间，从时间点January 1，1970 UTC到现在经历的时间（微秒）
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j+1] < array[j] {
				temp := array[j+1]
				array[j+1] = array[j]
				array[j] = temp
			}
		}
	}
	end := time.Now().UnixNano()
	sli := array[:]
	//将结果写入文件
	utils.StoreFiles(sli, "BubblingSortSmallestToLargest")
	return end - start
}

// 从大到小
func BubblingSortLargestToSmallest(array [80000]int) (ConstTime int64) { //从大到小排序
	start := time.Now().UnixNano()
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] < array[j+1] {
				temp := array[j+1]
				array[j+1] = array[j]
				array[j] = temp
			}
		}
	}
	end := time.Now().UnixNano()
	sli := array[:]
	//将结果写入文件
	utils.StoreFiles(sli, "BubblingSortLargestToSmallest")
	return end - start
}
