package main

import (
  "fmt"
)

func main() {
  // declare variable power of type int
  // by default this will be assigned value of 0
  var power int
  // Assign 9000 to power variable
  power = 9000
  fmt.Printf("It's over %d\n", power)

  // We can combine the two lines above
  // declare a variable number of type int and assign 100 to it
  var number int = 100
  fmt.Printf("It's over %d\n", number)

  // Both of these are not used very often
  // Go has a short cut for this:
  // declare variable and assign value to it, while inferring the type
  age := 26
  fmt.Printf("It's over %d\n", age)

  // Go only allows you to declare variables once, in the same scope.
  // so this will fail:
  // power := 9001
  // fmt.Printf("It's over %d\n", power)

  // Go also allows you to assign multiple variables
  name, power := "Goku", 9000
  fmt.Printf("%s's power is over %d\n", name, power)
}
