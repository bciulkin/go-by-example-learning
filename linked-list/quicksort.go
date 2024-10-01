package main

import (
  "cmp"
)

func partition[T cmp.Ordered](head *Node[T], tail *Node[T]) *Node[T] {
  pivot := head

  pre := head
  curr := head

  for curr != tail.next {
    if curr.val < pivot.val {
      temp := curr.val
      curr.val = pre.next.val
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

func quickSortHelper[T cmp.Ordered](head *Node[T], tail *Node[T]) {
  if (head.next == nil || head == tail) {
    return
  }

  // find pivot node
  pivot := partition(head, tail)

  // recursive call for less than pivot list
  quickSortHelper(head, pivot)

  // recursive call for greater than pivot list
  quickSortHelper(pivot.next, tail)
}

func QuickSort[T cmp.Ordered](head *Node[T]) {
  tail := tail(head)

  quickSortHelper(head, tail)
}

func tail[T cmp.Ordered](node *Node[T]) *Node[T] {
  tail := node
  for tail.next != nil {
    tail = tail.next
  }
  return tail

}
