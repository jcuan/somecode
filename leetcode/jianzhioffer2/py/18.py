# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def deleteNode(self, head: ListNode, val: int) -> ListNode:
        if head is None:
            return None
        preNode = None
        node = head
        while node is not None:
            if node.val == val:
                if preNode is None:
                    return node.next
                else:
                    preNode.next = node.next
                    return head
            preNode = node
            node = node.next
        return None