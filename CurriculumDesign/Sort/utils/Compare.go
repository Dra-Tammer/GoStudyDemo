package utils

import "fmt"

func Compare(nums ...int64) {
	var compareMap map[int64]string = map[int64]string{
		nums[0]: "bubbling",
		nums[1]: "shell",
		nums[2]: "insert",
		nums[3]: "select",
		nums[4]: "merge",
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j+1] < nums[j] {
				temp := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = temp
			}
		}
	}
	fmt.Printf("性能由高到低：")
	for _, v := range nums {
		fmt.Printf("\t%s\t", compareMap[v])
	}
	fmt.Printf("\n")
}
