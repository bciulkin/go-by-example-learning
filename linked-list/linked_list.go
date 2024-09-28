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

func (list *LinkedList[T]) Reverse() {
  // convert regular linked list to circular linked list
  circular := list.head
  length := 1
  for circular.next != nil {
    length++;
    circular = circular.next
  }
  circular.next = list.head

  // move x times (x = length)

  // un-circulate the linked list
  circular.next = nil
}

func (list *LinkedList[T]) Print() {
  var elements []T
  for e:= list.head; e != nil; e = e.next {
    elements = append(elements, e.val)
  }
  fmt.Println(elements)
}
