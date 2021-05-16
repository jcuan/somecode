from typing import List

class Solution:
    def exchange(self, nums: List[int]) -> List[int]:
        numcnt = len(nums)
        if numcnt <= 1:
            return nums
        s = 0
        end = numcnt - 1
        while s <= end:
            if nums[end] % 2 == 0:
                end -= 1
                continue
            if nums[s] % 2 == 1:
                s += 1
                continue
            nums[s], nums[end] = nums[end], nums[s]
            s += 1
            end -=1
        return nums