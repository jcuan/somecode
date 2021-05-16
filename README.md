# 刷题记录

leetcode还有看以前数据结构ppt遇到的

xxx0、xxx1分别代表同一问题的不同解法，后边数字大的一般更好，文件名代表leetcode题目链接编号

# think

## 通用

1. 不要僵化地实现算法
    - 涉及到反向等场景，想想中间容器是不是必须的，能否直接在原数据上操作或者直接保存结果：[jianzhioffer2/go/58.go](jianzhioffer2/go/58.go)
    - 分了很多个组一定要有什么group编号吗？不见得，通过step或者delta就行：[希尔排序](sort/insert.go) 看看生硬的`ShellSort_rubbish`
2. 涉及到值转换（在不同的标准下值不同）的情况下，不要误用
    - [jianzhioffer2/go/7.go](jianzhioffer2/go/7.go)中的`forBuildTree2`中的`valInorderIndexMap[root.Val]+1`，使用的是origin的index，但是误传入了进行了下标转换的`indexInorder`
3. 注意所有的循环遍历查找是否可以通过转换获得相关值
    - [jianzhioffer2/go/7.go](jianzhioffer2/go/7.go) 中的`buildTree1`->`buildTree2`
4. 分清递归需要调用的函数，[leetcode/jianzhioffer2/go/26.go](leetcode/jianzhioffer2/go/26.go), 在递归函数`forIsSubStructure0`出现了特殊的开始节点规避，是很stupid的做法，正确的是`forIsSubStructure`的实现，这个递归的处理，该在外层函数
5. 精简返回值:(
    - [leetcode/jianzhioffer2/py/54.py](leetcode/jianzhioffer2/py/54.py) inorder注释的那个实现可以说是非常垃圾了
    - go有些文章推荐使用int而不是uint本来就是因为负值可以代表很多的含义，[leetcode/jianzhioffer2/go/55.go](leetcode/jianzhioffer2/go/55.go)的`getBalancedInfo`最初还使用了一个error返回值来代表是否已经检测到失衡，实际上一个-1已经足够
6. 问题的进一步抽象，可以简化算法，增加可读性
    - [矩阵的顺序遍历](leetcode/jianzhioffer2/go/29.go) 开始没有引入top、down这种界限限制，导致下标操作可读性差，且存在后续的特殊处理

## 语言特定

1. go的stack，一般用动态数组即可，不需要使用list。`nonleetcode/go/simple_expression_cal.go`
2. 关于不要僵化地实现算法中，避免中间容器的使用，`jianzhioffer2/go/7.go::buildTree3`和`buildTree2`可以参考，不过可能大大牺牲了可读性，2->3主要是为了降低内存使用
3. copy拷贝的长度为src和dst的最小值


## 总结

1. 二叉树的递归和非递归遍历算法，都需要掌握，因为非递归方式保存了父节点的信息，很多题是或者可以是这种场景的变化
2. 二分法自己的原则：两端取等、中间偏左
    - `leetcode/jianzhioffer2/py/11.py`这道题错了非常多次，因为采用了`中间偏左`，即`mid = (start + end) // 2`的方式，所以有可能存在mid=end-1=start的情况，此时要收缩右部的区间，end=mid而不是end=mid-1！！

# 数据结构和算法汇总

## 排序

可以使用这个题目验证https://leetcode-cn.com/problems/sort-an-array/submissions/

- [插入排序：简单插入、折半插入、希尔排序](sort/insert.go) 
- [交换排序：冒泡、快排](sort/exchange.go)
- [归并排序](sort/merge.go)

## 回溯

非递归实现的效率贼低还贼复杂，非常不直观，调试了很久，难点主要在回退的处理

- [非递归实现，通过栈保存信息，处理相关回退](leetcode/jianzhioffer2/py/11.py)
- [递归实现, 由函数递推来自动处理下标变化和回退](leetcode/jianzhioffer2/go/11.go)

## 线性表

- [双栈应用：实现队列](https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/)
- [双栈应用：实现min函数的栈](https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/solution/mian-shi-ti-30-bao-han-minhan-shu-de-zhan-fu-zhu-z/)
- [链表：合并有序链表](structure/list.go) 新增一个冗余头节点后，可读性增加很多`MergeTwoIntList2`

## 查找和遍历

### 树

- [先序、中序、后序遍历：递归和非递归实现](structure/tree.go) 需要特别注意后序遍历的非递归实现
- [层次遍历](leetcode/jianzhioffer2/go/32.go) 敲重点！
- [使用先序和中序构造二叉树](leetcode/jianzhioffer2/go/7.go) [py](leetcode/jianzhioffer2/py/7.py)
    - 优化过程：优化循环查找->不使用个中间slice，只传下标
- [二叉搜索树：第K大结点](leetcode/jianzhioffer2/py/54.py)

重复的不太重要的：
- [判断树的子结构](leetcode/jianzhioffer2/go/26.go)
- [镜像二叉树](leetcode/jianzhioffer2/go/27.go)
- [检测是否是平衡二叉树、获取数的深度](leetcode/jianzhioffer2/go/55.go)
- [非递归后续遍历的应用：最近公共祖先](leetcode/jianzhioffer2/go/68.go)
- [二叉搜索树的后序遍历序列](leetcode/jianzhioffer2/py/33.py)

### 特殊查找

- [矩阵查找：特殊对角线](leetcode/jianzhioffer2/py/4.py)
- [二分查找](leetcode/jianzhioffer2/py/11.py) 需要特殊注意边界的处理，多1少1相关的错误

## 字符串

- [句子中单词顺序反转](leetcode/jianzhioffer2/go/58.go)
- [递归：重复字符串的排列](https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/solution/mian-shi-ti-38-zi-fu-chuan-de-pai-lie-hui-su-fa-by/)

## 动态规划

https://zhuanlan.zhihu.com/p/91582909

空间复杂度优化主要从状态转移方程入手

- [绳子切分后乘积](leetcode/jianzhioffer2/py/14.py) 注意这题初始化值不是开始值的解，需要特殊处理
- [二维数组动态规划](leetcode/0.go)  uniquePaths62，空间优化版本uniquePaths62_2
- [二维数组动态规划-状态转移方程涉及到求最值](leetcode/0.go) uniquePathsWithObstacles64
- [字符串编辑距离](https://zhuanlan.zhihu.com/p/91582909) 案例4，很重要的一点是

重复：
- [正则表达式匹配](https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof/solution/zhu-xing-xiang-xi-jiang-jie-you-qian-ru-shen-by-je/)

## 数学推导

- [基本不等式](leetcode/jianzhioffer2/py/14.py) 
- [求余数运算](https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/solution/mian-shi-ti-14-ii-jian-sheng-zi-iitan-xin-er-fen-f/) `(xy)⊙p=[(x⊙p)(y⊙p)]⊙p` ⊙ 代表取余
- [有限状态自动机](https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/solution/mian-shi-ti-20-biao-shi-shu-zhi-de-zi-fu-chuan-y-2/)
    - 定义所有状态列表
    - 定义状态转换矩阵stat，通过`p' = stat[p][x]`进行状态转换，如果不能进行转换，则失败
    - 定义合法结束状态
    - 最后的p是否在合法的结束状态
