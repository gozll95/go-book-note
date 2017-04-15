package stack

import "testing"

func TestPushPop(t *testing.T) {
	c := new(Stack)
	c.Push(5)
	if c.pop() != 5 {
		t.Log("Pop dosen't give 5")
		t.Fail()
	}
}
