package main

import (
  "fmt"
)

type Node[T any] struct {
  val T
  next *Node[T]
}

type LinkedList[T any] struct {
  head *Node[T]
}

func (list *LinkedList[T]) Add(value T) {
  nodeToAdd := &Node[T]{val: value, next: nil}
  if list.head == nil { // add first element
    list.head = nodeToAdd
  } else { // add non-first element
    x := list.head
    for x.next != nil {
      x = x.next
    }
    x.next = nodeToAdd
  }
}


// Reverse method uses stack to temporary hold all values
// Idea for improvement would be more robust implementation for a stack
func (list *LinkedList[T]) Reverse() LinkedList[T] {
  var stack []T
  tmp := list.head
  for tmp.next != nil {
    stack = append(stack, tmp.val)
    tmp = tmp.next
  }

  stack = append(stack, tmp.val)

  reverseList := LinkedList[T] {}
  for i := range len(stack) {
    reverseList.Add(stack[len(stack) - i - 1])
  }
  return reverseList
}

func (list *LinkedList[T]) Print() {
  var elements []T
  for e:= list.head; e != nil; e = e.next {
    elements = append(elements, e.val)
  }
  fmt.Println(elements)
}
