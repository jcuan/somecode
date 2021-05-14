from typing import List

class Solution:
    def verifyPostorder(self, postorder: List[int]) -> bool:
        def recur(start, rootidx): # 检查的区间, 这里边的值
            if start >= rootidx - 1:
                return True
            # 找到第一个>=postorder[rootidx]的
            for devideidx in range(start, rootidx + 1):
                if postorder[devideidx] >= postorder[rootidx]:
                    break
            for i in range (devideidx, rootidx): # 这区间的，必须不比root小
                if postorder[i] < postorder[rootidx]:
                    return False
            return recur(start, devideidx - 1) and recur(devideidx, rootidx-1)

        return recur(0, len(postorder)-1)