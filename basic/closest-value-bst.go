package basic

import (
	"algorithms/basic/shared"
	"fmt"
	"math"
)

func ClosestValueBstRecursive(node *shared.Node, closest int, num int) int {
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

func ClosestValueBstIterative(node *shared.Node, closest int, num int) int {
	for node != nil {
		valueAbs := int(math.Abs(float64(node.Value - num)))
		if valueAbs == 0 {
			return node.Value
		}
		if valueAbs < int(math.Abs(float64(closest-num))) {
			closest = node.Value
		}
		if num > node.Value {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	return closest
}

func GetClosestNode() *shared.Node {
	return &shared.Node{
		Value: 10,
		Left: &shared.Node{
			Value: 5,
			Left: &shared.Node{
				Value: 2,
				Left: &shared.Node{
					Value: 1,
				},
			},
			Right: &shared.Node{
				Value: 5,
			},
		},
		Right: &shared.Node{
			Value: 15,
			Left: &shared.Node{
				Value: 13,
				Right: &shared.Node{
					Value: 14,
				},
			},
			Right: &shared.Node{
				Value: 22,
			},
		},
	}
}

func DemoClosestValueBst() {
	node := GetClosestNode()
	fmt.Printf("[ClosestValueBstRecursive] %d -> %d\n", 12, ClosestValueBstRecursive(node, node.Value, 12))
	fmt.Printf("[ClosestValueBstIterative] %d -> %d\n", 12, ClosestValueBstIterative(node, node.Value, 12))
}
