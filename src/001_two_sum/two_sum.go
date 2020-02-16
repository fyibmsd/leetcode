package solution

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if j, ok := m[target-nums[i]]; ok {
			return []int{j, i}
		}

		m[nums[i]] = i
	}

	return []int{}
}
