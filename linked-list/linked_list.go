package main

import (
  "fmt"
)

type LinkedList[T any] struct {
  head, tail *element[T]
}

type element[T any] struct {
  next *element[T]
  val T
}

func (list *LinkedList[T]) Add(value T) {
  if list.tail == nil { // add first element
    list.head = &element[T]{val: value}
    list.tail = list.head // in one element list tail is at the same time a head
  } else { // add non-first element
    list.tail.next = &element[T]{val: value}
    list.tail = list.tail.next
  }
}

func (list *LinkedList[T]) Print() {
  var elements []T
  for e:= list.head; e != nil; e = e.next {
    elements = append(elements, e.val)
  }
  fmt.Println(elements)
}
