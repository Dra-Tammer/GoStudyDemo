package Menu

import (
	"fmt"
	"log"
)

func InitPrimaryMenu() {
	for true {
		var choice = 0
		fmt.Println("---------------------------------------")
		fmt.Println("请输入您的选择：1.排序\t2.数组\t3.退出")
		fmt.Println("---------------------------------------")
		_, err := fmt.Scan(&choice)
		if err != nil {
			log.Println(err)
		}
		switch choice {
		case 1:
			SortMenu()
		case 2:
			ArrayMenu()
		case 3:
			log.Println("退出系统")
			return
		default:
			log.Println("输入不合法")
		}
	}
}
