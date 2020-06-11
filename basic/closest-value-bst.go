package basic

import (
	"fmt"
	"math"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func ClosestValueBstRecursive(node *Node, closest int, num int) int {
	valueAbs := int(math.Abs(float64(node.Value - num)))
	if valueAbs == 0 {
		return node.Value
	}
	if valueAbs < int(math.Abs(float64(closest-num))) {
		closest = node.Value
	}
	if num > node.Value && node.Right != nil {
		return ClosestValueBstRecursive(node.Right, closest, num)
	} else if num < node.Value && node.Left != nil {
		return ClosestValueBstRecursive(node.Left, closest, num)
	}
	return closest
}

func DemoClosestValueBst() {
	node := Node{
		Value: 10,
		Left: &Node{
			Value: 5,
			Left: &Node{
				Value: 2,
				Left: &Node{
					Value: 1,
				},
			},
			Right: &Node{
				Value: 5,
			},
		},
		Right: &Node{
			Value: 15,
			Left: &Node{
				Value: 13,
				Right: &Node{
					Value: 14,
				},
			},
			Right: &Node{
				Value: 22,
			},
		},
	}
	fmt.Println(ClosestValueBstRecursive(&node, node.Value, 14))
}
