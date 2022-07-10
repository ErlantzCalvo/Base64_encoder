package main
import ("fmt"
          "strconv")

func main() {
  var a string
  fmt.Scanf("%s", &a)

  // Convert string to int array
  b := stringToASCIIArray(a)
  fmt.Println(b)

  // Convert each number of the array to binary
  fmt.Println("Binario:")
  c := intToBinaryArray(b)
  fmt.Println(c)
}

func stringToASCIIArray(str string) []int32 {
  return []rune(str)
}

func intToBinaryArray(arr []rune) []int32 {
  binArr := make([]int32, len(arr))
  for i:=0; i<len(arr); i++ {
    num, err := strconv.Atoi(fmt.Sprintf("%b", arr[i]))
    if err==nil {
      binArr[i] = int32(num)
    }
  }
  return binArr
}
