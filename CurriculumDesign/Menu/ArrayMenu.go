package Menu

import (
	"StudyDemo/CurriculumDesign/Array"
	"fmt"
	"log"
)

func ArrayMenu() {
	for true {
		var choice = 0
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("请输入您的选择:1.寻找马鞍点\t2.成绩统计\t3.退回上一级目录")
		fmt.Println("-----------------------------------------------------------------")
		_, err := fmt.Scan(&choice)
		if err != nil {
			log.Println(err)
		}
		switch choice {
		case 1:
			StartSaddlePoint()
		case 2:
			StartScoreStatistics()
		case 3:
			return
		default:
			log.Println("输入不合法")
		}
	}
}

func StartSaddlePoint() {
	SP := Array.SaddlePoint{}
	Array.SaddlePointHandle(SP.InitSaddlePointQuestion())
}

func StartScoreStatistics() {
	SCS := Array.ScoreStatistics{}
	Array.ScoreStatisticsHandle(SCS.InitScoreStatisticsQuestion())
}
