// Invert a binary tree.
// 
// Example:
// 
// Input:
// 
//      4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9
// Output:
// 
//      4
//    /   \
//   7     2
//  / \   / \
// 9   6 3   1
// 

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root!=nil{
        root.Left,root.Right = root.Right,root.Left
        invertTree(root.Left)
        invertTree(root.Right)
    }
    return root
}


