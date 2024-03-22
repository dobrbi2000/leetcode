package main

import "fmt"

//узел бинарного дерева
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//вставка узла
func insert(root *Node, key int) *Node {

	if root == nil {
		return &Node{Key: key}
	}

	if key == root.Key {
		fmt.Println("Дубликат ключа. Пропускаем вставку узла с ключом", key)
		return root
	}

	if key < root.Key {
		root.Left = insert(root.Left, key)
	} else if key > root.Key {
		root.Right = insert(root.Right, key)
	}

	return root

}

//обход бинарного дерева
func printInNode(root *Node) {
	if root != nil {
		printInNode(root.Left)
		fmt.Printf("%d ", root.Key)
		printInNode(root.Right)
	}
}

//минимальный элемент в дереве
func getMin(root *Node) int {
	if root.Left == nil {
		return root.Key
	}
	return getMin(root.Left)
}

func getMax(root *Node) int {
	if root.Right == nil {
		return root.Key
	}
	return getMin(root.Right)
}

func main() {
	root := &Node{Key: 10}
	insert(root, 5)
	insert(root, 14)
	insert(root, 15)
	insert(root, 8)
	insert(root, 16)
	insert(root, 3)
	insert(root, 22)
	insert(root, 10) // Дубликат ключа

	fmt.Println("Обход дерева:")
	printInNode(root)

	fmt.Println("\nМинимальный элемент:", getMin(root))
	fmt.Println("Максимальный элемент:", getMax(root))

}
