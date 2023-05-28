package functions

import (
	"StudyDemo/CurriculumDesign/Sort/utils"
	"time"
)

// ShellSort 希尔排序需要一个增量序列，在插入排序的时候保证序列尽量有序，把本身无需的数组变的部分有序，缩小增量排序
func ShellSort(array [80000]int) (ConstTime int64) {
	start := time.Now().UnixMilli()
	length := len(array)
	gap := 1
	//定义初始增量，也就是步长
	for gap < length/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := array[i] //记录需要插入的元素
			j := i - gap
			for j >= 0 && array[j] > temp { //把大的元素往右边插入
				array[j+gap] = array[j]
				j -= gap
			}
			array[j+gap] = temp
		}
		gap = gap / 3 //缩小增量
	}
	end := time.Now().UnixMilli()
	sli := array[:]
	utils.StoreFiles(sli, "ShellSort")
	return end - start
}
