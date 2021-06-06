package leetcode

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	}
	if nil == l2 {
		return l1
	}
	ret := &ListNode{0, nil}
	curListNode := ret
	var preKey int
	for nil != l1 || nil != l2 {
		var num1, num2 int
		if nil != l1 {
			num1 = l1.Val
			l1 = l1.Next
		}
		if nil != l2 {
			num2 = l2.Val
			l2 = l2.Next
		}
		add := num1 + num2 + preKey
		curListNode.Next = &ListNode{add % 10, nil}
		curListNode = curListNode.Next
		preKey = add / 10
	}
	if preKey != 0 {
		curListNode.Next = &ListNode{preKey, nil}
	}
	return ret.Next
}

func CountMap(m map[byte]int, key byte) bool {
	_, ok := m[key]
	return ok
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func LengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}
	n := len(s)
	left, right := 0, 0
	mapS := map[byte]int{}
	ret := 0
	for right < n {
		for CountMap(mapS, s[right]) {
			delete(mapS, s[left])
			left++
		}
		mapS[s[right]] = 1
		right++
		ret = max(ret, right-left)
	}
	return ret
}

func HammingDistance(x int, y int) int {
	temp := x ^ y
	countNum := func(num int) int {
		ret := 0
		for num != 0 {
			num = num & (num - 1)
			ret++
		}
		return ret
	}
	return countNum(temp)
}

func LongestPalindrome(s string) string {
	palindLen := func(s string, left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left++
			right--
		}
		return right - left - 1
	}
	maxLen := 0
	startIndex := 0
	for i := 0; i < len(s); i++ {
		leftLen, rightLen := palindLen(s, i, i), palindLen(s, i, i+1)
		if leftLen > rightLen {
			if leftLen > maxLen {
				maxLen = leftLen
				startIndex = i - maxLen/2
			}
		} else {
			if rightLen > maxLen {
				maxLen = rightLen
				startIndex = i - (maxLen-1)/2
			}
		}
	}
	return s[startIndex : startIndex+maxLen]
}

func TotalHammingDistance(nums []int) int {
	ret := 0
	for i := 0; i < 30; i++ {
		c := 0
		for _, val := range nums {
			c += (val >> i) & 1
		}
		ret += c * (len(nums) - c)
	}
	return ret
}
