package utils

import (
	"math/rand"
	"time"
)

func InitRand() (InitedArray [80000]int) {
	var MAXNUM = 80000
	var MINNUM = 20000
	var RandArray = make([]int, 1, 1)
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < MAXNUM; i++ {
		cot := rand.Intn(MAXNUM) - MINNUM
		RandArray = append(RandArray, cot)
	}
	var Array [80000]int
	for i := 0; i < 80000; i++ {
		Array[i] = RandArray[i]
	}
	return Array
}
