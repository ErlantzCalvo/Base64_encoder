package main

import (
        "fmt"
        "strings"
)

func main() {
  var a string
  fmt.Println("Insert a base64 encoded string:")
  fmt.Scanln(&a)

  if len(a) < 4 {
    panic("String cannot be shorter than 4 chars!")
  }
  numPaddings := CountPaddings(a)
  fmt.Println(numPaddings)

  fmt.Println("Get utf-8 char value of the string:")
  b := StringToUTF8(a)
  fmt.Println(b)

  fmt.Println("Get binary values of the string")
  c := Utf8ToBinary(b)
  fmt.Println(c)
}

func CountPaddings(str string) int8 {
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

func StringToUTF8(str string) []rune {
  return []rune(str)
}

func Utf8ToBinary(utfArr []rune) string {
  var result strings.Builder
  for i:=0; i<len(utfArr); i++ {
    binary:= fmt.Sprintf("%b", utfArr[i])
    result.WriteString(binary)
  }
  return result.String()
}
