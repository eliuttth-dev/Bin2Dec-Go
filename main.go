package main

import(
  "fmt"
  "strconv"
)

func main(){
  // Read input
  var binary string
  
  fmt.Print("Enter up to 8 binary digits (0's and 1's): ")
  fmt.Scanln(&binary)
  
  // Validate Input
  if !isValidBinary(binary){
    fmt.Println("Invalid input. Please enter only 0's and 1's.")
    return
  }

  // Convert binary to decimal
  decimal := binaryToDecimal(binary)

  // Display Output
  fmt.Printf("Decimal equivalent: %d\n", decimal);
}

// isValidBinary checks if the input consist only of 0's and 1's.
func isValidBinary(binary string) bool {
  for _, digit:= range binary{
    if digit != '0' && digit != '1' {
      return false
    }
  }
  return true
}

// binaryToDecimal converts a binary string to its decimal equivalent.
func binaryToDecimal(binary string) int {

  decimal := 0

  for i := len(binary) - 1; i >= 0; i--{
    bit, _ := strconv.Atoi(string(binary[i]))
    decimal += bit * pow(2,len(binary)-1-i)
  }
  return decimal
}

// pow calculates the power of base^exp.
func pow(base,exp int) int{
  
  result := 1
  for exp > 0 {
    if exp%2 == 1 {
      result *= base
    }
    base *= base
    exp /= 2
  }
  return result
}










