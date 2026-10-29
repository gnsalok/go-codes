package advancedgenerics

import "testing"

func TestStackUsesStaticTypes(t *testing.T) {
	var stack Stack[string]
	stack.Push("first")
	stack.Push("second")

	got, ok := stack.Pop()
	if !ok || got != "second" || stack.Len() != 1 {
		t.Fatalf("Pop() = %q, %v; Len() = %d", got, ok, stack.Len())
	}
}

func TestPopClearsRemovedValue(t *testing.T) {
	stack := Stack[*int]{}
	value := 42
	stack.Push(&value)

	_, _ = stack.Pop()
	backingArray := stack.items[:cap(stack.items)]
	if backingArray[0] != nil {
		t.Fatal("Pop retained a reference to the removed value")
	}
}
