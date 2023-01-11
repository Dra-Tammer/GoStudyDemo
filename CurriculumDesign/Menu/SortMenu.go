package Menu

import (
	"StudyDemo/CurriculumDesign/Sort/functions"
	"StudyDemo/CurriculumDesign/Sort/utils"
	"fmt"
	"log"
)

func SortMenu() {
	for true {
		var choice = 0
		fmt.Println("---------------------------------------------------------")
		fmt.Println("请输入您的选择：1.执行排序和比较程序\t2.返回上一级目录\t")
		fmt.Println("---------------------------------------------------------")
		_, err := fmt.Scan(&choice)
		if err != nil {
			log.Println(err)
		}
		switch choice {
		case 1:
			StartSorting()
		case 2:
			return
		default:
			log.Println("输入不合法")
		}
	}
}

func StartSorting() {
	Array := utils.InitRand()
	log.Println("数据初始化成功，随机生成了80000个数据")
	bsstlTime := functions.BubblingSort(Array)
	ssstlTime := functions.ShellSort(Array)
	isstlTime := functions.InsertSort(Array)
	sesstlTime := functions.SelectSort(Array)
	mergeArr := Array[:]
	mssTime := functions.MergeSort(mergeArr)
	fmt.Println("执行从大到小排序所花费的时间：")
	fmt.Printf("冒泡排序：%d毫秒\n", bsstlTime)
	fmt.Printf("希尔排序：%d毫秒\n", ssstlTime)
	fmt.Printf("插入排序：%d毫秒\n", isstlTime)
	fmt.Printf("选择排序：%d毫秒\n", sesstlTime)
	fmt.Printf("归并排序：%d毫秒\n", mssTime)
	utils.Compare(bsstlTime, ssstlTime, isstlTime, sesstlTime, mssTime)
}
