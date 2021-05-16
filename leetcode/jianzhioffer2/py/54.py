# Definition for a binary tree node.

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

# https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
class Solution:
    kk: int
    res: int

    def kthLargest(self, root: TreeNode, k: int) -> int:
        self.kk = k
        self.inorder2(root)
        return self.res
    
    def inorder2(self, root: TreeNode):
        if root is None:
            return None
        self.inorder2(root.right)
        if self.kk == 0:
            return
        self.kk -= 1
        if self.kk == 0:
            self.res = root.val
        self.inorder2(root.left)

    # # 返回非None代表已经找到
    # def inorder(self, root: TreeNode):
    #     if root is None:
    #         return None
    #     res = self.inorder(root.right)
    #     if res is not None:
    #         return res
    #     self.rank += 1
    #     if self.rank == self.k:
    #         return root.val
    #     res = self.inorder(root.left)
    #     if res is not None:
    #         return res

if __name__ == '__main__':
    root = TreeNode(5)
    root.left = TreeNode(3)
    root.right = TreeNode(6)
    root.left.left = TreeNode(2)
    root.left.right = TreeNode(4)
    root.left.left.left = TreeNode(1)

    Solution().kthLargest(root, 3)

