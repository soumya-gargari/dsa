package main

import "fmt"

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

func (n *node) delete(val int) *node {
	if n == nil {
		return n
	}
	if !n.search(val) {
		fmt.Println("value is not found")
		return n
	}
	if val < n.root {
		n.left = n.left.delete(val)
	} else if val > n.root {
		n.right = n.right.delete(val)
		// till here it is decided that which node to delete
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}
		// if it is reaching here, that means the node has both right and left child
		minVal := n.right.minVal()
		n.root = minVal
		n.right = n.right.delete(minVal)
	}
	return n
}

var predecessor int

func (n *node) inorderpredecessor(val int) int {
	if n == nil {
		return predecessor
	}
	if !n.search(val) {
		fmt.Println("value is not found")
		return predecessor
	}
	if n.root == val {
		if n.left != nil {
			return n.left.maxVal()
		}
	} else if val < n.root {
		return n.left.inorderpredecessor(val)
	} else if val > n.root {
		// if it is reaching here, that means provided value has no left child
		//if val > n.root {
		predecessor = n.root
		return n.right.inorderpredecessor(val)
	}

	return predecessor
}

var prev = &node{}

func isBST(n *node) bool {
	if n != nil {
		if !isBST(n.left) {
			return false
		}
		if prev != nil && prev.root >= n.root {
			return false
		}
		prev = n
		if !isBST(n.right) {
			return false
		}
	}
	return true
}

func (n *node) minVal() int {
	if n.left == nil {
		return n.root
	}
	return n.left.minVal()
}

func (n *node) maxVal() int {
	if n.right == nil {
		return n.root
	}
	return n.right.maxVal()
}

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
	fmt.Println("is it a valid binary tree:", isBST(bt))
	fmt.Println(bt.inorderpredecessor(20))
}
