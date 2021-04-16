package list

import "testing"

func TestBasic(t *testing.T) {
	l1 := New()
	l2 := New()

	l1.PushBack(10)
	l1.PushBack(11)
	l1.PushBack(14)
	l1.PushBack(13)

	l2.PushFront(1)
	l2.PushFront(2)
	l2.PushFront(3)
	l2.PushFront(5)
	l2.PushFront(4)

	l1.PushBackList(l2)
}
