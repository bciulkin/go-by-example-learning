package main

import (
  "fmt"
)

func main() {
  list := LinkedList[int] {}
  list.Add(4)
  list.Add(3)
  list.Add(1)
  list.Add(2)
  list.Add(5)

  fmt.Println("int list:")
  list.Print()

  tail := list.tail()
  fmt.Println(tail)

  reversedList := list.Reverse()
  fmt.Println("reversed list:")
  reversedList.Print()

  sortedList := list.quickSort()
  fmt.Println("sorted list:")
  sortedList.Print()

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
