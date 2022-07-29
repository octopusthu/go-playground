package leetcode

import "testing"

type findContentChildrenTestCase struct {
	expected int
	g        []int
	s        []int
}

var findContentChildrenTestCases = []findContentChildrenTestCase{
	{1, []int{1, 2, 3}, []int{1, 1}},
	{2, []int{1, 2}, []int{1, 2, 3}},
}

func TestFindContentChildren(t *testing.T) {
	for _, testCase := range findContentChildrenTestCases {
		if result := FindContentChildren(testCase.g, testCase.s); result != testCase.expected {
			t.Errorf("FindContentChildren(%v, %v) was incorrect, got: %d, want: %d.",
				testCase.g, testCase.s, result, testCase.expected)
		}
	}
}
