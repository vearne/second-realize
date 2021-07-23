package leecode

func LengthOfLongestSubstring(s string) int {
	// 左指针:i, 右指针:rp, 长度 rp -i + 1
	rp := -1
	ans := 0
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if i > 0 {
			delete(m, s[i-1])
		}
		for rp+1 < len(s) && m[s[rp+1]] == 0 {
			m[s[rp+1]]++
			rp++
		}
		ans = max(ans, rp-i+1)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
