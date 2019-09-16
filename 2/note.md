## 思路

新的链表，所有的节点需要新建而不是操作传入的参数

```
1.  p1、p2分别指向l1、l2
    resHead=new ListNode, 作为结果链表（第一个链表节点不使用，方便操作，最后去掉）, pres=resHead, pres.next=null
    carry=0用于标记是否有进位
2.  if p1 != null || p2 != null， 步骤3， 否则步骤6
3.  value = carry
    if p1 != null:
        value = value + p1.val; p1 = p1.next
    if p2 != null:
        value = value + p2.val; p2 = p2.next
    if value > 9:
        carry = 0
    else:
        carry = 1, value = value % 10
    newNode = new ListNode, newNode.val = value, newNode.next = null
    pres.next = NewNode, pres = pres.next
    去步骤2
5.  if carry!=0:
        newNode = new ListNode, newNode.val = carry, newNode.next = null
        pres.next = NewNode, pres = pre.next
6.  resHead = resHead.next （如果手动管理内存的，需要pres=resHead, free pres）
```