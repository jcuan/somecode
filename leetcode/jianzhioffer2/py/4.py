from typing import List

class Solution:
    def findNumberIn2DArray(self, matrix: List[List[int]], target: int) -> bool:
        if len(matrix) == 0 or len(matrix[0]) == 0 :
            return False
        rowMax = len(matrix) - 1
        columnMax = len(matrix[0]) - 1
        row = rowMax
        column = 0
        while 0<= row and column <= columnMax:
            nodeVal = matrix[row][column]
            if nodeVal == target:
                return True
            if nodeVal < target:
                column += 1
                continue
            if nodeVal > target:
                row -= 1
        return False

if __name__ == '__main__':
    print(Solution().findNumberIn2DArray([[-5]], -10))