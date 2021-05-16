from typing import List, Any, Mapping
from enum import IntEnum
from operator import add, attrgetter, sub

class Direction(IntEnum):
    UP = 0
    DOWN = 1
    LEFT = 2
    RIGHT = 3

class DirectionAction:
    def __init__(self, operatorFunc):
        self.operatorFunc = operatorFunc

class MatchedNode:
    def __init__(self, x: int, y: int, val: str):
        self.x = x
        self.y = y
        self.val = val

class Solution:
    xMax: int
    yMax: int
    x: int # 当前横坐标
    y: int # 当前纵坐标
    index: int # 当前匹配的字符的索引
    directionStack: List[Direction] # 注意保存的值，从A下移动到B，保存的是DOWN
    directionActions: List[DirectionAction]
    matchedNodes: List[MatchedNode]
    board: List[List[str]]

    def init(self, board):
        self.yMax = len(board[0]) - 1
        self.xMax = len(board) - 1
        self.directionStack = []
        self.matchedNodes = []
        self.board = board

        self.directionActions = [
            DirectionAction(lambda : (True, self.x - 1, self.y) if self.x > 0 else (False, None, None)),
            DirectionAction(lambda : (True, self.x + 1, self.y) if self.x < self.xMax else (False, None, None)),
            DirectionAction(lambda : (True, self.x, self.y - 1) if self.y > 0 else (False, None, None)),
            DirectionAction(lambda : (True, self.x, self.y + 1) if self.y < self.yMax else (False, None, None)),
        ]

    def setNextLocation(self, direction: Direction) -> bool:
        action = self.directionActions[direction]
        res, x, y = action.operatorFunc()
        if res == False:
            return False
        self.x = x
        self.y = y
        self.directionStack.append(direction)
        return True

    def nextLocation(self, matched):
        if matched: # 如果匹配，继续按照顺序来
            self.matchedNodes.append(MatchedNode(self.x, self.y, self.board[self.x][self.y]))
            self.board[self.x][self.y] = ''
            self.index += 1
            for direction in range(Direction.UP, Direction.RIGHT + 1):
                if self.setNextLocation(direction) == False:
                    continue
                return True
            # 需要回溯、所有的方向已经无法匹配
        # 需要回溯
        while len(self.directionStack) > 0: # 有可以回溯的
            lastDirection = self.directionStack.pop()
            # 还原回溯的坐标
            lastMatchedNode = self.matchedNodes[len(self.matchedNodes) - 1]
            self.x = lastMatchedNode.x
            self.y = lastMatchedNode.y
            if lastDirection != Direction.RIGHT: # 这个格子还有其他方向可匹配
                # 从下一个可能的方向开始
                for direction in range(lastDirection + 1, Direction.RIGHT + 1):
                    if self.setNextLocation(direction) == False:
                        continue
                    return True
            self.index -= 1 # 这个格子已经不可能匹配正确，正式回退
            self.matchedNodes.pop()
            self.board[self.x][self.y] = lastMatchedNode.val
        return False

    # 按照上、下、左、右的顺序遍历
    def exist(self, board: List[List[str]], word: str) -> bool:
        self.init(board)
        wordLen = len(word)
        def setStartPoint(x, y):
            self.x = x
            self.y = y
            self.index = 0
            self.directionStack.clear()
            self.matchedNodes.clear()
        for i in range (0, self.xMax + 1):
            for j in range(0, self.yMax + 1):
                setStartPoint(i, j)
                while self.index < wordLen:
                    if board[self.x][self.y] == word[self.index]:
                        if self.index == wordLen - 1:
                            return True
                        if self.nextLocation(True) == False: # 已经无法继续回溯，需要选择下一个中心点
                            break
                    else:
                        if self.nextLocation(False) == False: # 已经无法继续回溯，需要选择下一个中心点
                            break
        return False

if __name__ == '__main__':
    boards = [
        ["A","B","C","E"],
        ["S","F","C","S"],
        ["A","D","E","E"]
        # ['A', 'A', 'A', 'A'],
        # ['A', 'A', 'A', 'A'],
        # ['A', 'A', 'A', 'A'],
    ]
    data = {
        'ABCE': True,
        'S': True,
        'SADEC': True,
        # 'AAAAAAAAAAAAA':False
    }
    for s, r in data.items():
        if r != Solution().exist(boards, s):
            print('not equal', s, r)