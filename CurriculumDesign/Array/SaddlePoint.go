package Array

import (
	"fmt"
	"log"
)

type SaddlePoint struct {
	Array          [][]int
	TheSaddlePoint []int
	Row            int
	Col            int
}

func (SaddlePoint SaddlePoint) InitSaddlePointQuestion() (InitedArray SaddlePoint) {
	fmt.Println("请输入行数和列数：")
	_, err := fmt.Scan(&SaddlePoint.Row, &SaddlePoint.Col)
	if err != nil {
		log.Println("输入行数和列数这里", err)
	}
	SaddlePoint.Array = make([][]int, SaddlePoint.Row, SaddlePoint.Row)
	fmt.Println(fmt.Sprintf("请创建一个%d*%d的矩阵：", SaddlePoint.Row, SaddlePoint.Col))
	for i := 0; i < SaddlePoint.Row; i++ {
		for j := 0; j < SaddlePoint.Col; j++ {
			var Import int
			_, err := fmt.Scan(&Import)
			if err != nil {
				log.Println("矩阵的初始化这里", err)
			}
			SaddlePoint.Array[i] = append(SaddlePoint.Array[i], Import)
		}
	}
	fmt.Println(fmt.Sprintf("通过键盘输入创建了一个%d*%d的矩阵：", SaddlePoint.Row, SaddlePoint.Col))
	for _, rowValue := range SaddlePoint.Array {
		for _, colValue := range rowValue {
			fmt.Printf("%d\t", colValue)
		}
		fmt.Println()
	}
	return SaddlePoint
}

func SaddlePointHandle(point SaddlePoint) {
	for i := 0; i < point.Row; i++ {
		var min = point.Array[i][0]
		var minCol = 0
		for j := 0; j < point.Col; j++ {
			if point.Array[i][j] <= min {
				min = point.Array[i][j]
				minCol = j
			}
		}
		var max = point.Array[i][minCol]
		for k := 0; k < point.Row; k++ {
			if point.Array[k][minCol] > max {
				max = point.Array[k][minCol]
				break
			}
		}
		if max == point.Array[i][minCol] {
			point.TheSaddlePoint = append(point.TheSaddlePoint, max)
		}
	}
	if point.TheSaddlePoint == nil {
		fmt.Println("该矩阵没有马鞍点")
	} else {
		fmt.Println("马鞍点有：")
		for _, v := range point.TheSaddlePoint {
			fmt.Printf("%d\t", v)
		}
		fmt.Println()
	}
}
