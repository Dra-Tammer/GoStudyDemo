package utils

import (
	"testing"
)

// 写入文件操作的单元测试文件
func TestStoreFiles(t *testing.T) {
	arr := make([]int, 10, 10)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	StoreFiles(arr, "name")
	t.Logf("输出正确")
}
