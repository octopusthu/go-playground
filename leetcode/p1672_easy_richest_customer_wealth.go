package leetcode

func MaximumWealth(accounts [][]int) int {

	// A simple traversal

	var maximumWealth = 0
	for _, account := range accounts {
		var wealth = 0
		for _, value := range account {
			wealth += value
		}
		if wealth > maximumWealth {
			maximumWealth = wealth
		}
	}
	return maximumWealth
}
