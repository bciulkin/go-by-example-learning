package main

import (
  "cmp"
)

func binarySearch[T cmp.Ordered](head *Node[T], value T) bool {
  start := head
  var last *Node[T]
  last = nil

  for true {
    mid := middle(start, last)

    if mid == nil {
      return false
    }

    if mid.val == value {
      return true
    } else if start == last {
      break
    } else if mid.val < value {
      start = mid.next
    } else if mid.val > value {
      last = mid
    }
  }

  return false
}

func middle[T cmp.Ordered](start *Node[T], last *Node[T]) *Node[T] {
  if start == nil {
    return nil
  }

  if start == last {
    return start
  }

  slow := start
  fast := start.next

  for fast != last {
    fast = fast.next
    slow = slow.next
    if fast != last {
      fast = fast.next
    }
  }

  return slow
}
