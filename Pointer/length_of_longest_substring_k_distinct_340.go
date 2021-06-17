package Greedy

//给定一个字符串 s ，找出至多包含 k 个不同字符的最长子串 T。
//
//示例 1:
//
//输入: s = "eceba", k = 2
//输出: 3
//解释: 则 T 为 "ece"，所以长度为 3。

// lengthOfLongestSubstringKDistinct
// 可根据题655进行求解
// 用滑动窗口算法进行求解
// 1.判断特殊数据
// 2.右窗口滑动,将数据存储
// 3.判断存储的数据是否符合标准
// 4.不符合存储数据移动左窗口
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	//判断特殊数据
	if k==0{
		return len(s)
	}
	if len(s) ==0{
		return 0
	}

	l,max := 0,0
	//右窗口滑动
	for r := 0; r < len(s); r++ {
		s2 := s[l:r]
	}
	return 0
}