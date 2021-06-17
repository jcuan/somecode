# 错题集

## 为啥自己写的这么垃圾

- [链表公共节点](https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/submissions/)
    ```go
    for {
        if pa == nil {
            if !paChanged {
                pa = headB
                paChanged = true
            }
            // ...
        }
        if pb == nil {
            if !pbChanged {
                pb = headA
                pbChanged = true
            }
            // ...
        }
        // ...
        pa = pa.Next // warning！！！！这里可以这样做的前提是headB和headA都不能为空，否则这两行可能访问空指针
        pb = pb.Next
    }
    ```
    还有一点，看看人家写的。。。。
    ```go
    func getIntersectionNode(headA, headB *ListNode) *ListNode {
        curA, curB := headA, headB
        for curA != curB {
            if curA == nil {
                curA = headB
            } else {
                curA = curA.Next
            }
            if curB == nil {
                curB = headA
            } else {
                curB = curB.Next
            }
        }
        return curA
    }
    ```
- [有序数组的二分搜索](https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/)
    - [自己写的，代码多了两三倍吧](leetcode/jianzhioffer2/go/53.go)
    ```go
    func binaryFind(nums []int, target int) int {
        left := 0
        right := len(nums) - 1
        for left <= right {
            mid := (left + right) / 2
            if nums[mid] <= target {
                left = mid + 1
            } else {
                right = mid - 1 
            }
        }
        return left
    }

    func search(nums []int, target int) int {
        // 1. 排除明显不存在的
        if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
            return 0
        }
        // 分别找右边界，左边界
        return binaryFind(nums, target) - binaryFind(nums, target-1) 
    }
    ```
    两端不取等，可以同时处理下边的问题
    ```
    1. 找不到对应的target：
        序列：...,nums[i]=target-n,nums[j]=target+m,... (n>0, m >0)，
        考虑最终得到结果前，left可能指向i或者j，但是最终的都是一样的
        此时binaryFind左边得到的值是j，右边得到的值也是j，减出来就是0
    2. 能找到对应的target，且存在多个
        序列：...,nums[i-i]=target-n,nums[i]=target,...,nums[j-1]=target,nums[j]=target+m
        最终得到的结果：左边是i，右边是j，同样符合结果
    ```
    综上，核心其实就是：在使用了不取等的方法逼近区间后，可以使用target-1这种方式逼近区间左端点，同时处理等值区间（`[target,target,target]`）内的向右移动问题，提高代码复用度