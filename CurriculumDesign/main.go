package main

import "StudyDemo/CurriculumDesign/Array"

func main() {
	SP := Array.SaddlePoint{}
	Array.SaddlePointHandle(SP.InitSaddlePointQuestion())
	scs := Array.ScoreStatistics{}
	Array.ScoreStatisticsHandle(scs.InitScoreStatisticsQuestion())
}
