package leecode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	expectedMap := map[string]int{
		"abcaacbb": 3,
		" ":        1,
		"aaa":      1,
		"":         0,
		"pwwkew":   3,
	}
	for param, expeted := range expectedMap {
		got := LengthOfLongestSubstring(param)
		if expeted != got {
			t.Errorf("-LengthOfLongestSubstring, param: %v, expect: %v, got: %v", param, expeted, got)
		}
	}

}
