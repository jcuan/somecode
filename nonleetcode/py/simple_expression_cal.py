from typing import List, Mapping
from enum import Enum
from operator import add, sub, mul, ifloordiv
import math

class P(Enum):
    X = 0
    H = 1
    L = 2
    E = 3

operators = ['+', '-', '*', '/', '(', ')', '#']

priorityMatirx = [
    # +      -       *    /    (    )    #
    [P.H, P.H, P.L, P.L, P.L, P.H, P.H], # +
    [P.H, P.H, P.L, P.L, P.L, P.H, P.H], # -
    [P.H, P.H, P.H, P.H, P.L, P.H, P.H], # *
    [P.H, P.H, P.H, P.H, P.L, P.H, P.H], # /
    [P.L, P.L, P.L, P.L, P.L, P.E, P.X], # (
    [P.H, P.H, P.H, P.H, P.X, P.E, P.H], # )
    [P.L, P.L, P.L, P.L, P.L, P.X, P.E], # #
]

operatorIndexMap : Mapping[str, int] = {}

operatorFuncMap = {
    '+': add,
    '-': sub,
    '*': mul,
    '/': lambda x, y: int(x / y)
}

class ExpressionError(Exception):
    pass

def init():
    for index, opr in enumerate(operators):
        operatorIndexMap[opr] = index

def getOperatorPriority(one: str, two: str) -> P:
    idx1 = operatorIndexMap.get(one, -1)
    idx2 = operatorIndexMap.get(two, -1)
    if idx1 == -1 or idx2 == -1:
        errval = 'operator index not found, one[{}][{}] two[{}][{}]'.format(one, idx1, two, idx2)
        raise ExpressionError(errval)
    return priorityMatirx[idx1][idx2]

def isOperator(item: str) -> bool:
    return not (item >= '0' and item <= '9')

def calTopNums(operators: List[str], nums: List[int]):
    if len(nums) < 2:
        raise ExpressionError('calTopNums[1]bad expression')
    num1 = nums.pop()
    num2 = nums.pop()
    operator = operators.pop()
    if operator not in operatorFuncMap:
        raise ExpressionError('calTopNums[2]bad expression')
    nums.append(operatorFuncMap[operator](num2, num1))

def simpleExpressionCal(instr: str) -> int:
    init()
    operators: List[str] = []
    nums: List[int] = []
    instr = instr.replace(' ', '')
    instr = instr + '#'
    if instr[0] == '-' or instr[0] == '+':
        instr = "0" + instr
    instr.replace('(-', '(0-').replace('(+', '(0+')
    operators.append('#')
    idx = 0
    inLen = len(instr)
    while idx <  inLen:
        val = instr[idx]
        currentOperator = operators[len(operators) - 1]
        if isOperator(val):
            pri = getOperatorPriority(currentOperator, val)
            if pri == P.H:
                calTopNums(operators, nums)
                continue
            elif pri == P.L:
                operators.append(val)
            elif pri == P.E:
                if val == ')':
                    operators.pop()
                elif val == '#':
                    if len(nums) == 0:
                        raise ExpressionError('bad expression')
                    return nums.pop()
            else:
                raise Exception('unknown priority of operator: {} {}'.format(currentOperator, val))
            idx += 1
            continue
        # 不是operator
        startIndex = idx
        idx += 1
        while not isOperator(instr[idx]): # 一定不会越界
            idx += 1
        num = int(instr[startIndex: idx])
        nums.append(num)

if __name__ == '__main__':
    datas = [
        "",
        "1",
        "1+1-2+2*4+4/2",
        "11+11-22-10*10+20/20",
        "(0-3)/4",
        "5-3/2"
    ]
    res_rights = [0, 1, 10, -99, 0, 4]
    for i, expression in enumerate(datas):
        try:
            res = simpleExpressionCal(expression)
            res_right = res_rights[i]
        except ExpressionError as e:
            if expression != '':
                print('unexpected exception: {} {}'.format(expression, e))
            continue
        if res != res_right:
            print('not equal: {} {}'.format(res_right, res))