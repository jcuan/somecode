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
4. 分清递归需要调用的函数，`leetcode/jianzhioffer2/go/26.go`, 在递归函数`forIsSubStructure0`出现了特殊的开始节点规避，是很stupid的做法，正确的是`forIsSubStructure`的实现，这个递归的处理，该在外层函数
5. 精简返回值:(
    - `leetcode/jianzhioffer2/py/54.py` inorder注释的那个实现可以说是非常垃圾了
    - go有些文章推荐使用int而不是uint本来就是因为负值可以代表很多的含义，`leetcode/jianzhioffer2/go/55.go`的`getBalancedInfo`最初还使用了一个error返回值来代表是否已经检测到失衡，实际上一个-1已经足够

## 语言特定

1. go的stack，一般用动态数组即可，不需要使用list。`nonleetcode/go/simple_expression_cal.go`
2. 关于不要僵化地实现算法中，避免中间容器的使用，`jianzhioffer2/go/7.go::buildTree3`和`buildTree2`可以参考，不过可能大大牺牲了可读性，2->3主要是为了降低内存使用
3. copy拷贝的长度为src和dst的最小值

## 典型

- 二叉树遍历：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/solution/mian-shi-ti-54-er-cha-sou-suo-shu-de-di-k-da-jie-d/

## 总结

1. 二叉树的递归和非递归遍历算法，都需要掌握，因为非递归方式保存了父节点的信息，很多题是或者可以是这种场景的变化