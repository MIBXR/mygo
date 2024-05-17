package main

import "fmt"

// TreeNode 接口定义了二叉树节点需要实现的方法
type TreeNode interface {
    GetValue() interface{}          // 获取节点存储的值
    GetLeft() TreeNode              // 获取左子节点
    GetRight() TreeNode             // 获取右子节点
    SetLeft(node TreeNode)          // 设置左子节点
    SetRight(node TreeNode)         // 设置右子节点
}

// BinaryTree 接口定义了二叉树需要实现的方法
type BinaryTree interface {
    Insert(value interface{})       // 插入一个值
    Find(value interface{}) bool    // 查找一个值
    Remove(value interface{})       // 移除一个值
    Traverse() []interface{}        // 遍历二叉树
}

// Node 结构体实现了 TreeNode 接口
type Node struct {
    value interface{}
    left  *Node
    right *Node
}

// GetValue 获取节点值
func (n *Node) GetValue() interface{} {
    return n.value
}

// GetLeft 获取左子节点
func (n *Node) GetLeft() TreeNode {
    return n.left
}

// GetRight 获取右子节点
func (n *Node) GetRight() TreeNode {
    return n.right
}

// SetLeft 设置左子节点
func (n *Node) SetLeft(node TreeNode) {
    n.left = node.(*Node)
}

// SetRight 设置右子节点
func (n *Node) SetRight(node TreeNode) {
    n.right = node.(*Node)
}

// SimpleBinarySearchTree 实现 BinaryTree 接口
type SimpleBinarySearchTree struct {
    root *Node
}

// Insert 插入值
func (tree *SimpleBinarySearchTree) Insert(value interface{}) {
    newNode := &Node{value: value}
    if tree.root == nil {
        tree.root = newNode
    } else {
        insertNode(tree.root, newNode)
    }
}

// insertNode 递归辅助函数，用于插入新节点
func insertNode(current, newNode *Node) {
    if lessThan(newNode.value, current.value) {
        if current.left == nil {
            current.left = newNode
        } else {
            insertNode(current.left, newNode)
        }
    } else {
        if current.right == nil {
            current.right = newNode
        } else {
            insertNode(current.right, newNode)
        }
    }
}

// Find 查找值
func (tree *SimpleBinarySearchTree) Find(value interface{}) bool {
    return findNode(tree.root, value)
}

// findNode 递归辅助函数，用于查找值
func findNode(node *Node, value interface{}) bool {
    if node == nil {
        return false
    }
    if node.value == value {
        return true
    }
    if lessThan(value, node.value) {
        return findNode(node.left, value)
    }
    return findNode(node.right, value)
}

// Traverse 遍历二叉树
func (tree *SimpleBinarySearchTree) Traverse() []interface{} {
    var values []interface{}
    inorderTraversal(tree.root, &values)
    return values
}

// inorderTraversal 中序遍历
func inorderTraversal(node *Node, values *[]interface{}) {
    if node != nil {
        inorderTraversal(node.left, values)
        *values = append(*values, node.value)
        inorderTraversal(node.right, values)
    }
}

// lessThan 辅助函数，用于比较值
func lessThan(a, b interface{}) bool {
    return a.(int) < b.(int)
}

func main_tree() {
    tree := &SimpleBinarySearchTree{}
    tree.Insert(5)
    tree.Insert(3)
    tree.Insert(7)
    tree.Insert(1)
    tree.Insert(4)

    fmt.Println("Tree values:", tree.Traverse())
    fmt.Println("Find 7:", tree.Find(7))
    fmt.Println("Find 10:", tree.Find(10))
}

