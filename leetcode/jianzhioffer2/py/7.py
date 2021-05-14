from typing import Mapping, List

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    valInorderIndexMap: Mapping[int, int]
    preorder: List[int]
    inorder: List[int]

    def buildTree(self, preorder: List[int], inorder: List[int]) -> TreeNode:
        if len(preorder) == 0:
            return None
        self.preorder = preorder
        self.inorder = inorder
        self.valInorderIndexMap = {}
        for i, val in enumerate(inorder):
            self.valInorderIndexMap[val] = i
        return self.forBuildTree(0, 0, len(inorder))

    def forBuildTree(self, rootIndexInPreorder: int, left: int, right: int) -> TreeNode:
        node = TreeNode(self.preorder[rootIndexInPreorder])
        rootIndexInInorder = self.valInorderIndexMap[node.val]
        if left < rootIndexInInorder:
            node.left = self.forBuildTree(rootIndexInPreorder + 1, left, rootIndexInInorder)
        if rootIndexInInorder + 1 < right:
            node.right = self.forBuildTree(rootIndexInPreorder + (rootIndexInInorder - left) + 1, rootIndexInInorder + 1, right)
        return node

if __name__ == '__main__':
    Solution().buildTree([3,9,20,15,7], [9,3,15,20,7])