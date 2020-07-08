package basic

import (
	"algorithms/basic/shared"
	"reflect"
	"testing"
)

var bigNodeLimit = 25

func getBigNode() *shared.Node {
	node := &shared.Node{
		Value: 1,
	}
	insertLevel(node, 0, bigNodeLimit)
	return node
}

func insertLevel(node *shared.Node, i int, limit int) {
	if i < limit {
		node.Left = &shared.Node{
			Value: 2,
		}
		node.Right = &shared.Node{
			Value: 3,
		}
		i++
		insertLevel(node.Left, i, limit)
		insertLevel(node.Right, i, limit)
	}

}

func BenchmarkBranchSumRecursive(b *testing.B) {
	node := getBigNode()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := make([]int, 0)
		BranchSumRecursive(node, 0, &result)
	}
}

func BenchmarkBranchSumStack(b *testing.B) {
	node := getBigNode()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BranchSumStack(node, bigNodeLimit)
	}
}

func TestBranchSumRecursive(t *testing.T) {
	node := GetBranchSumNode()
	result := make([]int, 0)
	want := []int{15, 16, 18, 10, 11}
	if BranchSumRecursive(node, 0, &result); !reflect.DeepEqual(result, want) {
		t.Errorf("BranchSumRecursive() = %v, want %v", result, want)
	}
}

func TestBranchSumStack(t *testing.T) {
	node := GetBranchSumNode()
	want := []int{15, 16, 18, 10, 11}
	if got := BranchSumStack(node, 3); !reflect.DeepEqual(got, want) {
		t.Errorf("BranchSumStack() = %v, want %v", got, want)
	}
}
