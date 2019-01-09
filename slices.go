package main

import "fmt"

func main()  {
  scores := []int{1, 2, 3, 4, 5}
  slice := scores[2:4]
  slice[0] = 999
  fmt.Println(scores)
}
