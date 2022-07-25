package main
import ("fmt"
        "strconv"
        "strings"
        "bufio"
        "os"
      )

func main() {
  fmt.Println("Write a text to encode in Base64:")
  // fmt.Scanf("%s", &a)
  in := bufio.NewReader(os.Stdin)
  line, err := in.ReadString('\n')
  
  if err != nil {
    panic("Error ocurred")
  }
  
  line = strings.TrimSuffix(line,"\n")

  // Convert string to int array
  fmt.Println("Ascii numbers:")
  b := StringToASCIIArray(line)
  fmt.Println(b)

  // Convert each number of the array to binary
  fmt.Println("Binary:")
  c := IntToBinary(b)
  fmt.Println(c)

  // separate the full binary string to 6 digit groups
  fmt.Println("Split in groups of six:")
  d, padding := SeparateInSix(c)
  fmt.Println(d, padding)

  // convert new groups to decimal values
  fmt.Println("Convert new array to decimal values:")
  e := BinaryToDecimal(d)
  fmt.Println(e)

  // convert decimal values to corresponding chars
  fmt.Println("Convert array of decimal values to chars:")
  f := DecimalToChar(e)
  fmt.Println(f)

  // build final string encoded
  fmt.Println("Build the final encoded string:")
  base64String := BuildString(f, padding)
  fmt.Println(base64String)
}

func StringToASCIIArray(str string) []int32 {
  return []rune(str)
}

func IntToBinary(arr []rune) string {
  var binStr strings.Builder
  for i:=0; i<len(arr); i++ {
    binary := fmt.Sprintf("%b", arr[i])
    // Insert zeros until there are 8 digits
    for j:=0; j< (8 - len(binary)); j++ {
      binStr.WriteString("0")
    }
    binStr.WriteString(binary)
  }
  return binStr.String()
}

func SeparateInSix(str string) ([]string, int8) {
  arrLen := len(str) / 6
  carry := len(str) % 6
  var padding int8 = 0

  if carry > 0 {
    fillingZeros := 6 - carry
    arrLen ++
    for i:=0; i< fillingZeros; i++ {
      str += "0"
    }

    if fillingZeros == 2 {
      padding = 1
    } else {
      padding = 2
    }
  }

  var binArr = make([]string, arrLen)
  for i:=0; i< arrLen; i++ {
    currentPos := i*6
    binArr[i] = string(str[currentPos:currentPos+6])
  }
  return binArr, padding
}

func BinaryToDecimal(binArr []string) []int64 {
  decimalArr := make([]int64, len(binArr))
  
  for i:= 0; i< len(binArr); i++ {
    decimalValue, err := strconv.ParseInt(binArr[i], 2, 8)
    if err == nil {
      decimalArr[i] = decimalValue
    }
  }

  return decimalArr 
}

func DecimalToChar(decArray []int64) []byte {
  charArray := make([]byte, len(decArray))
  base64table := [64]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
                            'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
                            'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4','5', '6', '7',
                            '8', '9', '+', '/'}
  for i:=0; i<len(decArray); i++ {
    charArray[i] = base64table[decArray[i]]
  }

  return charArray
}

func BuildString(charArray []byte, padding int8) string {
  var result strings.Builder
  
  for i:=0; i<len(charArray); i++ {
    result.WriteByte(charArray[i])
  }
  
  var i int8
  for i= 0; i<padding; i++ {
    result.WriteByte('=')
  } 

  return result.String()
}
