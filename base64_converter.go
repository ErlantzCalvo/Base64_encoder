package main
import ("fmt"
        // "strconv"
        "strings"
      )

func main() {
  var a string
  fmt.Println("Write a text to encode in Base64:")
  fmt.Scanf("%s", &a)

  // Convert string to int array
  fmt.Println("Ascii numbers:")
  b := stringToASCIIArray(a)
  fmt.Println(b)

  // Convert each number of the array to binary
  fmt.Println("Binary:")
  c := intToBinary(b)
  fmt.Println(c)

  // separate the full binary string to 6 digit groups
  fmt.Println("Spearated in groups of six:")
  d := separateInSix(c)
  fmt.Println(d)
}

func stringToASCIIArray(str string) []int32 {
  return []rune(str)
}

func intToBinary(arr []rune) string {
  var binStr strings.Builder
  for i:=0; i<len(arr); i++ {
    binStr.WriteString(fmt.Sprintf("0%b", arr[i]))
  }
  return binStr.String()
}

func separateInSix(str string) []string {
  arrLen := len(str) / 6
  carry := 6 - (len(str) % 6)

  if carry > 0 {
    arrLen ++
    for i:=0; i< carry; i++ {
      str += "0"
    }
  }

  var binArr = make([]string, arrLen)
  for i:=0; i< arrLen; i++ {
    currentPos := i*6
    binArr[i] = string(str[currentPos:currentPos+6])
  }
  return binArr
}
