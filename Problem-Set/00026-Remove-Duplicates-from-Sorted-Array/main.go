package main

func removeDuplicates(nums []int) int {
	hash := map[int]bool{}

	for _, v := range nums {
		if _, ok := hash[v]; !ok {
			nums[len(hash)] = v
			hash[v] = true
		}
	}

	return len(hash)
}
