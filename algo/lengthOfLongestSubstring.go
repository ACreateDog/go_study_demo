package main

func lengthOfLongestSubstringV1(s string) int {
	if len(s) <= 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	var chMap = make(map[string]int)
	for i := range s {
		k := string(s[i])
		if _, ok := chMap[k]; !ok {
			chMap[k] = 0
		}
		chMap[k]++
	}
	maxSubLength := len(chMap)
	minSubLength := 1
	max := 1
	step := 1
	for i := 0; i < len(s); i += step {
		var tmp = make(map[string]int)
		for window := minSubLength; window <= maxSubLength; window++ {
			idx := i + window - 1
			if idx >= len(s) {
				i = len(s)
				break
			}
			k := string(s[idx])
			if _, ok := tmp[k]; ok {
				step = tmp[k] + 1 - i
				break
			}
			tmp[k] = idx
		}
		if max < len(tmp) {
			max = len(tmp)
		}
	}

	return max
}
func lengthOfLongestSubstring(s string) int { //双指针滑动窗口
	max := 0
	n := len(s)
	i := 0
	hashmap := make(map[string]int)
	for j := 0; j < n; j++ {
		hashmap[string(s[j])]++
		for hashmap[string(s[j])] > 1 {
			i++
			hashmap[string(s[i])]--
		}
		if j-i+1 > max {
			max = j - i + 1
		}
	}
	return max
}