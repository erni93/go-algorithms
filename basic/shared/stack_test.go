package shared

import (
	"reflect"
	"testing"
)

func getTestStackItem(value int) BranchSumItem {
	return BranchSumItem{
		Node: &Node{
			Value: value,
		},
		Sum: 0,
	}
}

func TestStack_Pop(t *testing.T) {
	fullStack := NewStack(2)
	fullStack.Push(getTestStackItem(1))
	fullStack.Push(getTestStackItem(2))
	emptyStack := NewStack(1)
	t.Run("Pop full stack", func(t *testing.T) {
		want := getTestStackItem(2)
		if got := fullStack.Pop(); !reflect.DeepEqual(got, want) {
			t.Errorf("Pop() = %v, want %v", got, want)
		}
	})
	t.Run("Pop empty stack", func(t *testing.T) {
		defer func() { _ = recover() }()
		emptyStack.Pop()
		t.Errorf("Pop() did not panic")
	})
}

func TestStack_Push(t *testing.T) {
	stack := NewStack(1)
	t.Run("Push a new item", func(t *testing.T) {
		want := getTestStackItem(2)
		stack.Push(want)
		if got := stack.items[0]; !reflect.DeepEqual(got, want) {
			t.Errorf("Push() = %v, want %v", got, want)
		}
	})
}
