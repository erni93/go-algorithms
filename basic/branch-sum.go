package basic

import (
	"algorithms/basic/shared"
	"fmt"
)

func BranchSumRecursive(node *shared.Node, sum int, sums *[]int) {
	nodeSum := node.Value + sum
	if node.Left == nil && node.Right == nil {
		*sums = append(*sums, nodeSum)
	}
	if node.Left != nil {
		BranchSumRecursive(node.Left, nodeSum, sums)
	}
	if node.Right != nil {
		BranchSumRecursive(node.Right, nodeSum, sums)
	}
}

func BranchSumStack(node *shared.Node, depth int) []int {
	sums := make([]int, 0, 2^depth)
	if node.Left == nil && node.Right == nil {
		return sums
	}
	stack := shared.NewStack(depth)
	sum := node.Value
	active := node.Left
	stackRightNode(node, stack, sum)
	for active != nil {
		sum += active.Value
		stackRightNode(active, stack, sum)
		if active.Left == nil && active.Right == nil {
			sums = append(sums, sum)
			if stack.Size > 0 {
				stackItem := stack.Pop()
				active = stackItem.Node
				sum = stackItem.Sum
			} else {
				break
			}
		} else {
			active = active.Left
		}
	}
	return sums
}

func stackRightNode(node *shared.Node, stack *shared.Stack, sum int) {
	if node.Left != nil && node.Right != nil {
		stack.Push(shared.BranchSumItem{
			Node: node.Right,
			Sum:  sum,
		})
	}
}

func GetBranchSumNode() *shared.Node {
	return &shared.Node{
		Value: 1,
		Left: &shared.Node{
			Value: 2,
			Left: &shared.Node{
				Value: 4,
				Left: &shared.Node{
					Value: 8,
				},
				Right: &shared.Node{
					Value: 9,
				},
			},
			Right: &shared.Node{
				Value: 5,
				Left: &shared.Node{
					Value: 10,
				},
			},
		},
		Right: &shared.Node{
			Value: 3,
			Left: &shared.Node{
				Value: 6,
			},
			Right: &shared.Node{
				Value: 7,
			},
		},
	}
}

func DemoBranchSum() {
	node := GetBranchSumNode()
	result := make([]int, 0)
	BranchSumRecursive(node, 0, &result)
	fmt.Printf("[BranchSumRecursive] -> %v\n", result)
	fmt.Printf("[BranchSumStack] -> %v\n", BranchSumStack(node, 3))
}
