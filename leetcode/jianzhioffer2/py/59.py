from collections import deque


# https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/
class MaxQueue:

    def __init__(self):
        self.q = deque()
        self.max_q = deque() # 保存所有的

    def max_value(self) -> int:
        if len(self.max_q) > 0:
            return self.max_q[0]
        return -1
    

    def push_back(self, value: int) -> None:
        self.q.append(value)
        while len(self.max_q) > 0 and self.max_q[-1] < value:
            self.max_q.pop()
        self.max_q.append(value)

    def pop_front(self) -> int:
        if len(self.q) == 0:
            return -1
        val  = self.q.popleft()
        if val == self.max_q[0]:
            self.max_q.popleft()
        return val

if __name__ == '__main__':
    obj = MaxQueue()
    obj.push_back(1)
    assert(obj.max_value()==1)
    obj.push_back(3)
    obj.push_back(4)
    assert(obj.max_value()==4)
    assert(obj.pop_front()==1)
    assert(obj.pop_front()==3)
    obj.push_back(2)
    assert(obj.max_value()==4)
    assert(obj.pop_front()==4)
    assert(obj.max_value()==2)
    obj.push_back(5)
    assert(obj.max_value()==5)