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
