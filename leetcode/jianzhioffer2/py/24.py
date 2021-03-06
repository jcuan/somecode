# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        if head is None or head.next is None:
            return head
        newHead = head
        remain = head.next
        head.next = None
        while remain is not None:
            tmp = remain
            remain = remain.next
            tmp.next = newHead
            newHead = tmp
        return newHead