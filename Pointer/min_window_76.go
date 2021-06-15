package Pointer

import "strings"

//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//示例 1：
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"

// minWindow 判断最短子串出现 t
// 滑动窗口算法
// 1.判断长度
// 3.先判断能判断的最小字符串
// 4.存储index,lastIndex
// 5.依次往右，重复操作
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	if len(s) == len(t) && s == t {
		return t
	}
	index := strings.IndexAny(s, t)
	lastIndex := strings.LastIndexAny(s, t)
	if index == -1 || lastIndex == -1 {
		return ""
	}
	tmpS := s[index : lastIndex+1]
	check := true
	//从left开始一直持续
	for check == true {
		check = false
		//将第一个删去,找到最近的
		if index = strings.IndexAny(tmpS[1:], t); index < 0 {
			break
		}

		//判断式子还是否成立
		for i := 0; i < len(t); i++ {
			//判断最新的临时数据是否成立
			if strings.Contains(tmpS[1+index:], string(t[i])) {
				if i == len(t)-1 && len(t)<=len(tmpS[1+index:]) {
					check = true
					tmpS = tmpS[1+index:]
				}
				continue
			}
			break
		}
	}

	check = true
	//从right开始一直持续
	for check == true {
		check = false
		//将最后一个删去,找到最近的
		if lastIndex = strings.LastIndexAny(tmpS[:len(tmpS)-1], t); lastIndex < 0 {
			break
		}
		//判断式子还是否成立
		for i := 0; i < len(t); i++ {
			//判断最新的临时数据是否成立
			if strings.Contains(tmpS[:lastIndex+1], string(t[i])) {
				if i == len(t)-1 && len(t)<=len(tmpS[:lastIndex+1]) {
					check = true
					tmpS = tmpS[:lastIndex+1]
				}
				continue
			}
			break
		}
	}

	return tmpS
}
