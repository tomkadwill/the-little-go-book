package main

import (
  "fmt"
)

func main() {
  value, exists := power("goku")
  fmt.Println("value: ", value)
  fmt.Println("exists: ", exists)

  _, exists2 := power("goku")
  fmt.Println("exists2: ", exists2)

  fmt.Println("1 + 2 = ", add(1, 2))
}

func log(message string)  {

}

func add(a, b int) int {
  return a + b
}

func power(name string) (int, bool) {
  return 1, true
}
