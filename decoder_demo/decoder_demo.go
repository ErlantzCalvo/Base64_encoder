package main

import (
        "fmt"
)

func main() {
  var a string
  fmt.Println("Insert a base64 encoded string:")
  fmt.Scanln(&a)

  if len(a) < 4 {
    panic("String cannot be shorter than 4 chars!")
  }
  numPaddings := countPaddings(a)
  fmt.Println((numPaddings))
}

func countPaddings(str string) int8 {
  var paddings int8 = 0
  for i:=0; i<2; i++ {
    if str[len(str)-1-i] == 61 {
      paddings++
    } else {
      break
    }
  }
  return paddings
}
