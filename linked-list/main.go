package main

import (
  "fmt"
)

func main() {
  list := LinkedList[int] {}
  list.Add(4)
  list.Add(3)
  list.Add(1)
  list.Add(5)
  list.Add(2)

  fmt.Println("int list:")
  list.Print()

  reversedList := list.Reverse()
  fmt.Println("reversed list:")
  reversedList.Print()

  QuickSort(list.head)
  fmt.Println("sorted list:")
  list.Print()

  fmt.Println(binarySearch(list.head, 3)) // true
  fmt.Println(binarySearch(list.head, 6)) // false
  fmt.Println(binarySearch(list.head, 2)) // true
  // *********************************
  /*stringList := LinkedList[string] {}
  stringList.Add("raz")
  stringList.Add("dwa")

  fmt.Println("string list:")
  stringList.Print()
*/
  // *********************************
/*
  fooList := LinkedList[foo] {}

  fmt.Println("foo list:")
  fooList.Print()
  */
}

type foo struct {
  bar int
  text string
}
