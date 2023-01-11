package Array

import (
	"fmt"
	"log"
	"math/rand"
)

type ScoreStatistics struct {
	StuNum   int
	ScoreArr []int
}

func (ScoreStatistics ScoreStatistics) InitScoreStatisticsQuestion() (InitedScoreStatistics ScoreStatistics) {
	fmt.Println("请输入参加计算机考试的学生人数：")
	_, err := fmt.Scan(&ScoreStatistics.StuNum)
	if err != nil {
		log.Println(err)
	}
	ScoreStatistics.ScoreArr = make([]int, ScoreStatistics.StuNum, ScoreStatistics.StuNum)
	for i := 0; i < ScoreStatistics.StuNum; i++ {
		ScoreStatistics.ScoreArr[i] = rand.Intn(100)
	}
	fmt.Printf("随机生成了%d人的数据，最终的统计结果如下:", ScoreStatistics.StuNum)
	return ScoreStatistics
}

func ScoreStatisticsHandle(scs ScoreStatistics) {
	var Statatistics = make(map[int]int)
	for i := 0; i < scs.StuNum; i++ {
		Statatistics[scs.ScoreArr[i]]++
	}
	var assist = 0
	for k, v := range Statatistics {
		if v > 1 {
			if assist%7 == 0 {
				fmt.Printf("\n")
			}
			assist++
			fmt.Printf("%d出现了%d次\t", k, v)
		}
	}
	fmt.Printf("\n")
}
