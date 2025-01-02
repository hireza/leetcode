package main

func removeElement(nums []int, val int) int {
	res := []int{}

	for _, v := range nums {
		if v == val {
			continue
		}

		res = append(res, v)
	}

	copy(nums, res)

	return len(res)
}
