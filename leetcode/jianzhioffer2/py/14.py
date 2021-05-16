import math

# https://leetcode-cn.com/problems/jian-sheng-zi-lcof/ 
# 由数学推导可知道约接近3越大
class Solution:
    def cuttingRope2(self, n: int) -> int:
        if n <= 3:
            return n - 1
        zhishu = n // 3
        remain = n % 3
        if remain == 0:
            return int(math.pow(3, zhishu))
        if remain == 1:
            return int(math.pow(3, zhishu-1) * (2 * 2))
        return int(math.pow(3, zhishu) * 2)

    def cuttingRope(self, n: int) -> int:
        if n <= 3:
            return n - 1
        dp = [0, 1, 2, 3] # 前边两个没有实际意义
        for i in range(4, n + 1):
            turns = i // 2
            maxVal = 0
            for j in range(1, turns + 1):
                value = dp[j] * dp[i-j]
                if value >  maxVal:
                    maxVal = value
            dp.append(maxVal)
        return dp[n]

if __name__ == '__main__':
    print(Solution().cuttingRope(9))
    print(Solution().cuttingRope(8))