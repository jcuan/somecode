from typing import List

class Solution:
    def minArray(self, numbers: List[int]) -> int:
        start = 0
        end = len(numbers) - 1
        while start < end:
            mid = (start + end) // 2
             # 当不能整除的时候，选择靠左的这一个，会>=start
            if numbers[mid] > numbers[end]:
                start = mid + 1
            elif numbers[mid] < numbers[end]:
                end = mid # 不能+1，因为有可能是[...10,1,2,...]这种情况，start->1,end->2,mid->1
            else:
                end -= 1
        return numbers[start]

if __name__ == '__main__':
    datas = [
        [3,1,3]
    ]
    for data in datas:
        print(Solution().minArray(data))