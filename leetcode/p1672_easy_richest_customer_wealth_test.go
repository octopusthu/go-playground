package leetcode

import "testing"

type maximumWealthTestCase struct {
	expected int
	accounts [][]int
}

var maximumWealthTestCases = []maximumWealthTestCase{
	{6, [][]int{{1, 2, 3}, {3, 2, 1}}},
	{10, [][]int{{1, 5}, {7, 3}, {3, 5}}},
	{17, [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}},
}

func TestMaximumWealth(t *testing.T) {
	for _, testCase := range maximumWealthTestCases {
		if result := MaximumWealth(testCase.accounts); result != testCase.expected {
			t.Errorf("MaximumWealth(%v) was incorrect, got: %d, want: %d.",
				testCase.accounts, result, testCase.expected)
		}
	}
}
