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
	_, err := fmt.Scanf("%d", &ScoreStatistics.StuNum)
	if err != nil {
		log.Println(err)
	}
	ScoreStatistics.ScoreArr = make([]int, ScoreStatistics.StuNum, ScoreStatistics.StuNum)
	for i := 0; i < ScoreStatistics.StuNum; i++ {
		ScoreStatistics.ScoreArr[i] = rand.Intn(100)
	}
	fmt.Println(ScoreStatistics.ScoreArr)
	return ScoreStatistics
}

func ScoreStatisticsHandle(scs ScoreStatistics) {
	var Statatistics = make(map[int]int)
	for i := 0; i < scs.StuNum; i++ {
		Statatistics[scs.ScoreArr[i]]++
	}
	for k, v := range Statatistics {
		if v > 1 {
			fmt.Printf("%d出现了%d次\n", k, v)
		} else {
			fmt.Println("成绩没有值重复出现")
		}
	}
}
