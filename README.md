# 刷题记录

leetcode还有看以前数据结构ppt遇到的

# think

## 通用

1. 不要僵化地实现算法
    - 涉及到反向等场景，想想中间容器是不是必须的，能否直接在原数据上操作或者直接保存结果：`jianzhioffer2/go/58.go`
2. 涉及到值转换（在不同的标准下值不同）的情况下，不要误用
    - `jianzhioffer2/go/7.go`中的forBuildTree2中的`valInorderIndexMap[root.Val]+1`，使用的是origin的index，但是误传入了进行了下标转换的`indexInorder`
3. 注意所有的循环遍历查找是否可以通过转换获得相关值
    - `jianzhioffer2/go/7.go::buildTree1`->`buildTree2`

## 语言特定

1. go的stack，一般用动态数组即可，不需要使用list。`nonleetcode/go/simple_expression_cal.go`
2. 关于不要僵化地实现算法中，避免中间容器的使用，`jianzhioffer2/go/7.go::buildTree3`和`buildTree2`可以参考，不过可能大大牺牲了可读性，2->3主要是为了降低内存使用
