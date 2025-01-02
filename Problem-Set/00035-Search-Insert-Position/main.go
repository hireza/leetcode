package main

func searchInsert(nums []int, target int) int {
	res := 0
	for i, v := range nums {
		if target > v {
			res = i + 1
			continue
		}

		if target == v {
			res = i
		}

		return res
	}

	return res
}
