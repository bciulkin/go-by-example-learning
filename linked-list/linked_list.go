package main

import (
  "fmt"
  "cmp"
)

type Node[T cmp.Ordered] struct {
  val T
  next *Node[T]
}

type LinkedList[T cmp.Ordered] struct {
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

func (list *LinkedList[T]) partition(head *Node[T], tail *Node[T]) *Node[T] {
  pivot := head

  fmt.Println("pivot: ", pivot)

  pre := head
  curr := head

  for curr.next != tail.next {
    if curr.val < pivot.val {
      fmt.Println("curr.val: ", curr.val)
      fmt.Println("pivot.val: ", pivot.val)
      temp := curr.val
      fmt.Println("temp: ", temp)
      curr.val = pre.next.val
      fmt.Println("pre.next.val: ", pre.next.val)
      pre.next.val = temp

      // Move pre to next node
      pre = pre.next
    }

    // move curr to next node
    curr = curr.next
  }

  // swap pivot data with pre data
  currValue := pivot.val
  pivot.val = pre.val
  pre.val = currValue

  return pre
}

func (list *LinkedList[T]) quickSortHelper(head *Node[T], tail *Node[T]) {
  if (head.next == nil || head == tail) {
    return
  }

  // find pivot node
  pivot := list.partition(head, tail)

  // recursive call for less than pivot list
  list.quickSortHelper(head, pivot)

  // recursive call for greater than pivot list
  list.quickSortHelper(pivot.next, tail)
}

func (list *LinkedList[T]) quickSort() *LinkedList[T] {
  list.quickSortHelper(list.head, list.tail())
  return list
}

func (list *LinkedList[T]) tail() *Node[T] {
  tail := list.head
  for tail.next != nil {
    tail = tail.next
  }
  return tail

}

func (list *LinkedList[T]) binarySearch() {
  
}

func (list *LinkedList[T]) Print() {
  var elements []T
  for e:= list.head; e != nil; e = e.next {
    elements = append(elements, e.val)
  }
  fmt.Println(elements)
}
