package functions

import (
	"StudyDemo/CurriculumDesign/Sort/utils"
	"time"
)

// BubblingSort 冒泡，和相邻的比较，一次循环之后，最大值就出现在最右边
func BubblingSort(array [80000]int) (ConstTime int64) { //从小到大排序
	start := time.Now().UnixMilli() //获取当前的时间，从时间点January 1，1970 UTC到现在经历的时间（微秒）
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j+1] < array[j] {
				temp := array[j+1]
				array[j+1] = array[j]
				array[j] = temp
			}
		}
	}
	end := time.Now().UnixMilli()
	sli := array[:]
	utils.StoreFiles(sli, "BubblingSort")
	return end - start
}
