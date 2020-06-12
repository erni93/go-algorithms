package basic

import (
	"fmt"
	"testing"
)

var closestNumbers = [4]int{5, 3, 12, 18}

func BenchmarkClosestValueBstRecursive(b *testing.B) {
	node := GetNode()
	for _, num := range closestNumbers {
		b.Run(fmt.Sprintf("n=%d", num), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ClosestValueBstRecursive(&node, node.Value, num)
			}
		})
	}
}

func BenchmarkClosestValueBstIterative(b *testing.B) {
	node := GetNode()
	for _, num := range closestNumbers {
		b.Run(fmt.Sprintf("n=%d", num), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ClosestValueBstIterative(&node, node.Value, num)
			}
		})
	}
}
