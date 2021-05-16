class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        plen = len(p)
        slen = len(s)
        sindex = slen - 1
        pindex = plen - 1
        while pindex >= 0 and pindex >= 0:
            if p[pindex] == '.':
                pindex -= 1
                sindex -= 1
            elif p[pindex] == '*':
                # pindex此时一定大于0，否则是个错误的表达式
                char = p[pindex - 1]
                if char != s[sindex]:
                    if pindex == 1 and sindex == -1
                    pindex -= 2
            else:
                if s[sindex] != p[pindex]:
                    return False
                if pindex == 0 and sindex == 0:
                    return True
                pindex -= 1
                sindex -= 1