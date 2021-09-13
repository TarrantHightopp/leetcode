package main

import "fmt"

func main() {
	nums := []int{3, 3, 3, 5, 5, 5}
	target := 8
	fmt.Println(TowSunV1(nums, target))
}

// 结论值 = 指定值 - 元素值
// 检查map中是否有结论值
// 返回map下表值和当前loop的值
func TowSunV1(data []int, target int) []int {
	if len(data) <= 0 {
		return nil
	}
	var m = make(map[int]int)
	for i := 0; i < len(data); i++ {
		another := target - data[i]

		if _, ok := m[another]; ok {
			return []int{m[data[i]], i}
		}

		// 使用map 记录每个元素下表
		m[data[i]] = i
	}
	return nil
}
