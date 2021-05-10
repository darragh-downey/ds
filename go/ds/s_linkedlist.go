package ds

import "sync"

type S_Element struct {
	next  *S_Element
	Value interface{}
	lock  *sync.RWMutex
}

type S_List struct {
	sentinel S_Element
	len      int
	lock     sync.RWMutex
}

func (l *S_List) prepend(value interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	new_node := &S_Element{Value: value}

	new_node.next = &l.sentinel
	l.sentinel = *new_node
}

func (l *S_List) insertAfter(prev *S_Element, value interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if prev == nil {
		return
	}

	new_node := S_Element{Value: value}
	// make next of new element as next of prev element
	new_node.next = prev.next
	// make next of prev element equal new_node
	prev.next = &new_node
}

func (l *S_List) append(value interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	new_node := S_Element{Value: value}

	if l.sentinel == (S_Element{}) {
		l.sentinel = new_node
	}

	last := l.sentinel

	for last.next != nil {
		last = *last.next
	}
	last.next = &new_node
}
