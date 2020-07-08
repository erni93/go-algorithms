package shared

type Stack struct {
	items []BranchSumItem
	Size  int
}

type BranchSumItem struct {
	Node *Node
	Sum  int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		items: make([]BranchSumItem, capacity),
		Size:  0,
	}
}

func (stack *Stack) Push(data BranchSumItem) {
	stack.items[stack.Size] = data
	stack.Size++
}

func (stack *Stack) Pop() BranchSumItem {
	if stack.Size == 0 {
		panic("Pop from empty stack")
	}
	popValue := stack.items[stack.Size-1]
	stack.Size--
	return popValue
}
