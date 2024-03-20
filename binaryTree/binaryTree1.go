package main

import "fmt"

type node struct {
	root  int
	left  *node
	right *node
}

func (n *node) insert(val int) {
	if n.root == 0 {
		n.root = val
		return
	}
	if val < n.root {
		// goto left side
		if n.left == nil {
			n.left = &node{root: val}
		} else {
			n.left.insert(val)
		}
	} else if val > n.root {
		// go to right side
		if n.right == nil {
			n.right = &node{root: val}
		} else {
			n.right.insert(val)
		}
	}
}

func (n *node) search(val int) bool {
	if n == nil {
		return false
	}
	if n.root == val {
		return true
	}
	if val < n.root {
		// goto left side
		return n.left.search(val)

	} else if val > n.root {
		// go to right side
		return n.right.search(val)
	}
	return true
}

func (n *node) delete(val int) *node {
	if n == nil {
		return n
	}
	if !n.search(val) {
		fmt.Println("value is not found, not able to delete")
		return n
	}
	if val < n.root {
		// goto left side
		n.left = n.left.delete(val)

	} else if val > n.root {
		// go to right side
		n.right = n.right.delete(val)
		// till above node is decided which one to delete
	} else {

		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}
		// if it is reaching here, that means the value we want to delete, have both right and left children
		minVal := n.right.minVal()
		n.root = minVal
		n.right = n.right.delete(minVal)
	}
	return n
}

var predecessor = 0

func (n *node) inorderpredecessor(val int) int {
	if n == nil {
		return 0
	}
	if !n.search(val) {
		fmt.Println("value is not found, not able to find inorderpredecessor")
		return 0
	}

	if n.root == val {
		if n.left != nil {
			return n.left.maxVal()
		}
	} else if val < n.root {
		// goto left side
		return n.left.inorderpredecessor(val)

	} else {
		if val > n.root {
			// if it is reachinmg here , that means it dont have any left childs
			// so return just the previous value of that node
			// go to right side
			predecessor = n.root
			return n.right.inorderpredecessor(val)
			// till above node is decided which one to delete
		}
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
		if isBST(n.right) {
			return true
		}
	}
	return true
}

func displayLeftNodes(n *node) {
	if n == nil {
		return
	}
	fmt.Println("left node val is:", n.root)
	displayLeftNodes(n.left)
}

func displayRightNodes(n *node) {
	if n == nil {
		return
	}
	fmt.Println("Right node val is:", n.root)
	displayLeftNodes(n.right)
}

func (n *node) minVal() int {
	if n.left == nil {
		fmt.Println("min val in bst is:", n.root)
		return n.root
	}
	return n.left.minVal()
}

func (n *node) maxVal() int {
	if n.right == nil {
		fmt.Println("max val in bst is:", n.root)
		return n.root
	}
	return n.right.maxVal()
}

func main() {
	bst := &node{root: 5}
	bst.insert(2)
	bst.insert(8)
	bst.insert(1)
	bst.insert(9)
	displayLeftNodes(bst.left)
	displayRightNodes(bst.right)
	bst.minVal()
	bst.maxVal()
	fmt.Println("is value found:", bst.search(10))
	fmt.Println("is it a bst or not:", isBST(bst))
	bst.delete(8)
	displayRightNodes(bst.right)
}
