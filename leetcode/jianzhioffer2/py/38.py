from typing import List

# 原作者 https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/solution/mian-shi-ti-38-zi-fu-chuan-de-pai-lie-hui-su-fa-by/
class Solution:
    def permutation(self, s: str) -> List[str]:
        ss, res = list(s), [] # ss主要是为了便于交换，s不可修改
        slen = len(s)
        def dfs(fixedidx: int):
            if fixedidx == slen - 1:
                res.append(''.join(ss))
            dic = {}
            # fixedidx 当前固定的位置
            for i in range(fixedidx, slen):
                if ss[i] in dic:
                    continue
                dic[ss[i]] = True
                ss[i], ss[fixedidx] = ss[fixedidx], ss[i]
                dfs(fixedidx + 1)
                ss[i], ss[fixedidx] = ss[fixedidx], ss[i]
        dfs(0)
        return res

if __name__ == '__main__':
    print(Solution().permutation('abba'))