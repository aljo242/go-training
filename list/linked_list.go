package list

type Element struct {
	next, prev *Element
	data       int
	list       *List
}

type List struct {
	root Element
	len  int
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}

	return nil
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}

	return nil
}

func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

// Front returns the first element of list l or nil if empty
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}

	return l.root.next
}

// Front returns the last element of list l or nil if empty
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}

	return l.root.prev
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert inserts element e after element at, increments len and returns e
func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// convenience wrapper for insert
func (l *List) insertValue(data int, at *Element) *Element {
	return l.insert(&Element{data: data}, at)
}

func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--
	return e
}

// move moves e to next to at and returns e
func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}

	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e

	return e
}

// Remove removes e from l if e is an element of list l
// it returns the element value e.Value
// the elemetn must not be nil
func (l *List) Remove(e *Element) int {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil
		l.remove(e)
	}

	return e.data
}

// PushFront inserts a new element e with value data at the front of list l and returns e
func (l *List) PushFront(data int) *Element {
	l.lazyInit()
	return l.insertValue(data, &l.root)
}

// PushBack inserts a new element e with value data at the back of list l and returns e
func (l *List) PushBack(data int) *Element {
	l.lazyInit()
	return l.insertValue(data, l.root.prev)
}

// InsertBefore inserts a new element e with value data immediately before mark and returns e
func (l *List) InsertBefore(data int, mark *Element) *Element {
	if mark.list != l {
		return nil
	}

	return l.insertValue(data, mark.prev)
}

// InsertAfter inserts a new element e with value data immediately after mark and returns e
func (l *List) InsertAfter(data int, mark *Element) *Element {
	if mark.list != l {
		return nil
	}

	return l.insertValue(data, mark)
}

// MoveToFront moves element e to the front of list l
// if e is not an element of l, the list is not modified
// the element must not be niul
func (l *List) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}

	l.move(e, &l.root)
}

// MoveToBack moves element e to the back of list l
// if e is not an element of l, the list is not modified
// the element must not be niul
func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}

	l.move(e, l.root.prev)
}

// MoveBefore moves element e to its new position before mark
// if e or mark is not an element of l, or e == mark, list is unmodified
func (l *List) MoveBefore(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}

	l.move(e, mark.prev)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is unmodified
func (l *List) MoveAfter(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}

	l.move(e, mark)
}

// PushBackList inserts a copy of another list at the back of list l.
// the lists l and other may be the same
func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.data, l.root.prev)
	}
}

// PushFrontList inserts a copy of another list at the front of list l.
// the lists l and other may be the same
func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.data, &l.root)
	}
}
