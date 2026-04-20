package main
import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

func create_add(root* TreeNode, new_value int) *TreeNode{
	if (root == nil) {
		return &TreeNode{Value: new_value};
	}

	if root.Value < new_value {
		root.Right = create_add(root.Right, new_value);
	} else if (root.Value > new_value) {
		root.Left = create_add(root.Left, new_value);
	}

	return root;
}

func sum_of_tree(root *TreeNode) int{
	if (root == nil) {
		return 0;
	}
	sum := root.Value;
	sum += sum_of_tree(root.Left);
	sum += sum_of_tree(root.Right);
	return sum;
}


func main() {
	root := &TreeNode{Value: 5};
	root = create_add(root, 3);
	root = create_add(root, 2);
	root = create_add(root, 1);
	root = create_add(root, 4);
	root = create_add(root, 8);
	root = create_add(root, 7);
	root = create_add(root, 20);
	root = create_add(root, 22);
	fmt.Println(sum_of_tree(root));
}