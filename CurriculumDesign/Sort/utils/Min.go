package utils

func Min(nums ...int64) (minNum int64) {
	var m = nums[0]
	for i := 0; i < len(nums); i++ {
		if m > nums[i] {
			m = nums[i]
		}
	}
	return m
}
