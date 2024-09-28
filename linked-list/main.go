package main

import (
  "fmt"
)

func main() {
  list := LinkedList[int] {}
  list.Add(1)
  list.Add(2)
  list.Add(3)
  list.Add(4)

  //list.Reverse()

  fmt.Println("int list:")
  list.Print()

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
