package leetcode

import "sort"

func FindContentChildren(g []int, s []int) int {

	// 1. Sort g[] and s[] in ascending order
	// 2. For each g[i] find the smallest satisfying s[j]
	// 3. Stop when no satisfying s[j] could be found for g[i]

	sort.Ints(g)
	sort.Ints(s)
	var contentChildren = 0
	var i, j = 0, 0
	for i < len(g) {
		for j < len(s) && g[i] > s[j] {
			j++
		}
		if j < len(s) {
			contentChildren++
			j++
		} else if j == len(s) {
			break
		}
		i++
	}
	return contentChildren
}
