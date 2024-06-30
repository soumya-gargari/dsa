// Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.

// A leaf is a node with no children.

// Example 1:

// Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
// Output: true
// Explanation: The root-to-leaf path with the target sum is shown.
// Example 2:

// Input: root = [1,2,3], targetSum = 5
// Output: false
// Explanation: There two root-to-leaf paths in the tree:
// (1 --> 2): The sum is 3.
// (1 --> 3): The sum is 4.
// There is no root-to-leaf path with sum = 5.
// Example 3:

// Input: root = [], targetSum = 0
// Output: false
// Explanation: Since the tree is empty, there are no root-to-leaf paths.

package main

import (
	"fmt"
)

type node struct {
	left  *node
	right *node
	root  int
}

func (n *node) insert(val int) {
	if n.root == 0 {
		n.root = val
		return
	}
	if val < n.root {
		if n.left == nil {
			newNode := &node{root: val}
			n.left = newNode
		} else {
			n.left.insert(val)
		}
	} else if val > n.root {
		if n.right == nil {
			newNode := &node{root: val}
			n.right = newNode
		} else {
			n.right.insert(val)
		}
	}
}

// func (n *node) minVal() int {
// 	if n.left == nil {
// 		return n.root
// 	}
// 	return n.left.minVal()
// }

// func (n *node) maxVal() int {
// 	if n.right == nil {
// 		return n.root
// 	}
// 	return n.right.maxVal()
// }

func (n *node) search(val int) bool {
	if n == nil {
		return false
	}
	if val < n.root {
		return n.left.search(val)
	} else if val > n.root {
		return n.right.search(val)
	}
	return true
}

func displayLeftNodes(n *node) {
	if n == nil {
		return
	}
	fmt.Println("left node value is:", n.root)
	displayLeftNodes(n.left)
}

func displayRightNodes(n *node) {
	if n == nil {
		return
	}
	fmt.Println("right node value is:", n.root)
	displayRightNodes(n.right)
}

//Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf
//path such that adding up all the values along the path equals targetSum.

//A leaf is a node with no children.

// func hasPathSum(n *node, targetSum int) bool {
// 	var dfs func(n *node, currentSum int) bool
// 	dfs = func(n *node, currentSum int) bool {
// 		if n == nil {
// 			return false
// 		}
// 		currentSum += n.root
// 		if currentSum == targetSum && n.left == nil && n.right == nil {
// 			return true
// 		}
// 		return dfs(n.left, currentSum) || dfs(n.right, currentSum)
// 	}
// 	return dfs(n, 0)
// }

func main() {
	bt := &node{root: 10}
	bt.insert(5)
	bt.insert(20)
	bt.insert(1)
	bt.insert(6)
	bt.insert(15)
	bt.insert(30)
	displayLeftNodes(bt.left)
	displayRightNodes(bt.right)
	fmt.Println(bt.search(15))
	fmt.Println("is path val or target val found:", hasPathSum(bt, 16))
}

func hasPathSum(n *node, targetSum int) bool {
	var dfs func(n *node, currentSum int) bool
	dfs = func(n *node, currentSum int) bool {
		if n == nil {
			return false
		}
		currentSum += n.root
		if currentSum == targetSum && n.left == nil && n.right == nil {
			return true
		}
		return dfs(n.left, currentSum) || dfs(n.right, currentSum)
	}
	return dfs(n, 0)
}
