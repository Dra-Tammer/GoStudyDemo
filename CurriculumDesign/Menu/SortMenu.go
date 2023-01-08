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
		fmt.Println("请输入您的选择：1.执行排序和比较程序(从小到大)\t2.执行排序和比较程序(从大到小)\t3.返回上一级目录")
		_, err := fmt.Scan(&choice)
		if err != nil {
			log.Println(err)
		}
		switch choice {
		case 1:
			StartSortingSmallestToLargest()
		case 2:
			StartSortingLargestToSmallest()
		case 3:
			return
		default:
			log.Println("输入不合法")
		}
	}
}

func StartSortingSmallestToLargest() {
	Array := utils.InitRand()
	log.Println("数据初始化成功，随机生成了80000个数据")
	bsstlTime := functions.BubblingSortSmallestToLargest(Array)
	ssstlTime := functions.ShellSortSmallestToLargest(Array)
	isstlTime := functions.InsertSortSmallestToLargest(Array)
	sesstlTime := functions.SelectSortSmallestToLargest(Array)
	fmt.Println("执行从大到小排序所花费的时间：")
	fmt.Printf("冒泡排序：%d纳秒\n", bsstlTime)
	fmt.Printf("希尔排序：%d纳秒\n", ssstlTime)
	fmt.Printf("插入排序：%d纳秒\n", isstlTime)
	fmt.Printf("选择排序：%d纳秒\n", sesstlTime)
	minNum := utils.Min(bsstlTime, ssstlTime, isstlTime, sesstlTime)
	switch minNum {
	case bsstlTime:
		fmt.Println("可见希尔排序的效率最高")
	case ssstlTime:
		fmt.Println("可见希尔排序的效率最高")
	case isstlTime:
		fmt.Println("可见插入排序的效率最高")
	case sesstlTime:
		fmt.Println("可见选择排序的效率最高")
	default:
		return
	}
}

func StartSortingLargestToSmallest() {
	Array := utils.InitRand()
	log.Println("数据初始化成功，随机生成了80000个数据")
	bsltsTime := functions.BubblingSortLargestToSmallest(Array)
	ssltsTime := functions.ShellSortLargestToSmallest(Array)
	isltsTime := functions.InsertSortLargestToSmallest(Array)
	sesltsTime := functions.SelectSortLargestToSmallest(Array)
	fmt.Println("执行从大到小排序所花费的时间：")
	fmt.Printf("冒泡排序：%d纳秒\n", bsltsTime)
	fmt.Printf("希尔排序：%d纳秒\n", ssltsTime)
	fmt.Printf("插入排序：%d纳秒\n", isltsTime)
	fmt.Printf("选择排序：%d纳秒\n", sesltsTime)
	minNum := utils.Min(bsltsTime, ssltsTime, isltsTime, sesltsTime)
	switch minNum {
	case bsltsTime:
		fmt.Println("可见冒泡排序的效率最高")
	case ssltsTime:
		fmt.Println("可见希尔排序的效率最高")
	case isltsTime:
		fmt.Println("可见插入排序的效率最高")
	case sesltsTime:
		fmt.Println("可见选择排序的效率最高")
	default:
		return
	}
}
