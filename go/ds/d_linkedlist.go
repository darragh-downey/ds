package ds

import "sync"

// seeking to add RW protection to https://go.googlesource.com/go/+/go1.16.3/src/container/list/list.go#48
// for demonstration purposes only
type D_Element struct {
	prev, next *D_Element

	list *D_List

	Value interface{}

	lock *sync.RWMutex
}

func (e *D_Element) Prev() *D_Element {
	e.lock.Lock()
	defer e.lock.Unlock()
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *D_Element) Next() *D_Element {
	e.lock.Lock()
	defer e.lock.Unlock()
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// D_List is a doubly-linked list based on https://go.googlesource.com/go/+/go1.16.3/src/container/list/list.go#48
type D_List struct {
	root D_Element     // sentinel list element, only &root, root.prev and root.next are used
	len  int           // current length of the list
	lock *sync.RWMutex // RWrite protect len
}

func New() *D_List { return new(D_List).Init() }

func (l *D_List) Init() *D_List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func (l *D_List) Len() int {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.len
}

func (l *D_List) Front() *D_Element {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *D_List) Back() *D_Element {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *D_List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e
func (l *D_List) insert(e, at *D_Element) *D_Element {
	l.lock.Lock()
	defer l.lock.Unlock()
	e.prev = at
	e.next = at.next
	e.prev.next = e // equivalent of at.next = e
	e.next.prev = e // equivalent of at.next.prev = e
	e.list = l
	l.len++
	return e
}
