package binarytrees

import (
	"fmt"
	"log"
)

type Node struct {
	data   int
	left   *Node
	right  *Node
	parent *Node
}

func (node *Node) GetData() int {
	return node.data
}

func (node *Node) SetData(data int) {
	node.data = data
}

type BinaryTree struct {
	Root Node
}

func (btre *BinaryTree) SetRootData(data int) {
	btre.Root.data = data
}

func (btree *BinaryTree) Traverse() {
	recursiveTraverse(&btree.Root)

}
func (btree *BinaryTree) Add(newNode *Node) {
	addDj(&btree.Root, newNode)
}

func addDj(c_root, newNode *Node) {
	queue := [][2]*Node{{c_root, c_root}}
	var currentNode *Node
	for len(queue) != 0 {
		currentNode = queue[0][0]
		if currentNode == nil {
			//parent of this node is where we insert
			parent := queue[0][1]
			if parent.left == nil {
				parent.left = newNode
			} else if parent.right == nil {
				parent.right = newNode
			}
			return
		}
		nodeAndparent := [2]*Node{currentNode.left, currentNode}
		queue = append(queue, nodeAndparent)
		nodeAndparent[0] = currentNode.right
		queue = append(queue, nodeAndparent)
		queue = queue[1:]
	}
	log.Fatal("ERROR::INSERTING TO BTREE::Unable to insert")
}

func recursiveTraverse(c_root *Node) {
	queue := []*Node{c_root}
	var currentNode *Node
	for len(queue) != 0 {
		currentNode = queue[0]
		fmt.Println(currentNode.data)
		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}
		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
		queue = queue[1:]
	}
}
