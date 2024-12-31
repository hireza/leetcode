package main

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}

	for i, v := range nums {
		if idx, ok := hash[v]; ok {
			return []int{idx, i}
		}

		hash[target-v] = i
	}

	return nil
}
