from typing import List


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def getKthFromEnd(self, head: ListNode, k: int) -> ListNode:
        idx = 0
        node = head
        knode = head
        while node != None:
            if idx >= k:
                knode = knode.next
            node = node.next
            idx += 1
        return knode

if __name__ == '__main__':
    head = ListNode(0)
    head.next = ListNode(1)
    head.next.next = ListNode(2)
    head.next.next.next = ListNode(3)
    head.next.next.next.next = ListNode(4)

    print(Solution().getKthFromEnd(head, 2).val)