package jzoffer

// https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof/submissions/

/*
 dp[i][j] 代表i=len(s)，j=len(p)
 注意dp的[i][j]代表的是s[i-1]，p[j-1]，因为dp[0][0]代表了空串和空串的匹配
*/
func isMatch(s string, p string) bool {
	slen := len(s)
	plen := len(p)
	dp := make([][]bool, slen+1)
	for i := 0; i <= slen; i++ {
		dp[i] = make([]bool, plen+1)
	}
	dp[0][0] = true // 都为空时，可以匹配
	for i := 1; i <= slen; i++ {
		dp[i][0] = false // s不为空，p为空，明显不能匹配
	}
	// 当s为空，p不为空时，要复杂一点，可能出现x*x*x*x*这种情况
	for i := 1; i <= plen; i++ {
		if i%2 == 0 && p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		} else {
			dp[0][i] = false
		}
	}
	var nexti int
	var nextj int
	for i := 1; i <= slen; i++ {
		for j := 1; j <= plen; j++ {
			switch {
			// 假设字符串为"xc*"
			case p[j-1] == '*':
				nextj = j - 2  // 考察"xc*"前边的那个x
				if nextj < 0 { // *前边没有东西，表达式错误，直接返回false
					return false
				}
				if dp[i][nextj] { // c*不参与匹配成功
					dp[i][j] = true
					continue // end
				}
				// c*参与匹配，存在.*作为p开头的情况
				c := p[j-2]
				nexti = i // 从当前的开始,考察
				matched := false
				for nexti >= 1 && (s[nexti-1] == c || c == '.') {
					if dp[nexti-1][nextj] { // 根据前边的逻辑，nextj作为下标没有问题
						dp[i][j] = true
						matched = true
						break
					}
					nexti--
				}
				// 说明已经不可能产生匹配
				dp[i][j] = matched
			case p[j-1] == '.': // 取决于各去掉一个字符的匹配情况
				dp[i][j] = dp[i-1][j-1]
			default: // 取决于各去掉一个字符的匹配情况+本地的匹配情况
				dp[i][j] = dp[i-1][j-1] && s[i-1] == p[j-1]
			}
		}
	}
	return dp[slen][plen]
}

// -------warong answer

// 当前s[sidx]等于p[pidx]的时候，向前移动sidx和pidx
func nextIndex(s string, p string, sidx, pidx int) (int, int) {
	nexts := sidx - 1
	for nexts >= 0 && s[nexts] == s[sidx] {
		nexts--
	}
	nexts++
	nextp := pidx - 1
	for nextp >= 0 && p[nextp] == p[pidx] {
		nextp--
	}
	nextp++
	return nexts, nextp
}

/*
	错误的方法：无法处理下边的情况`s="aaa" p="ab*a*c*a"`，因为这种情况如果不使用动态规划，是要进行回溯的
	但是又因为*的原因，回溯的逻辑会比较复杂，因为有可能回溯到很远的地方
	dp[i][j]： i->s index, j-> p index
*/
func isMatchWrong(s string, p string) bool {
	i := len(s) - 1
	j := len(p) - 1
	var nextj int
	var nexti int
	for i >= 0 && j >= 0 {
		switch {
		case p[j] == '*':
			nextj = j - 1
			if nextj < 0 { // 正则格式错误
				return false
			}
			if p[nextj] == '.' { // 最复杂的情况
				dotIndex := nextj
				nextj = j - 2                       // 判断前边的是否相等
				if nextj >= 0 && s[i] == p[nextj] { // 假设此时相等的字符为c
					// 查看s中有多少个相等的
					nexti, nextj = nextIndex(s, p, i, nextj)
					// p中含有c个数(dotIndex-nextj)，如果> s总含有的个数，匹配失败
					if dotIndex-nextj > i-nexti+1 {
						return false
					}
					i = nexti - 1
					j = nextj - 1
				} else {
					// 不相等的情况, 或者.剩的最后一个字符(此时nextj<0)，这个.必须匹配，还是需要找s中有多少个相等的字符串
					// 查看s中有多少个相等的
					nexti = i - 1
					for nexti >= 0 && s[nexti] == s[i] {
						nexti--
					}
					nexti++ // 多走了一个
					i = nexti - 1
					j = dotIndex - 1
				}
			} else { // 不是*并且不是.
				if p[nextj] != s[i] { // 不相等的时候，p的c*直接作废
					j = nextj - 1
					continue
				}
				nexti, nextj = nextIndex(s, p, i, nextj)
				// p中含有c个数((j-1)-nextj + 1)减去1，如果> s总含有的个数，匹配失败
				if ((j-1)-nextj+1)-1 > i-nexti+1 {
					return false
				}
				i = nexti - 1
				j = nextj - 1
			}
		case p[j] == '.':
			i--
			j--
		default:
			if s[i] != p[j] {
				return false
			}
			i--
			j--
		}
	}
	return i == j || (j == 1 && p[j] == '*')
}
